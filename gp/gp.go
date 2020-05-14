package gp

import (
	"fmt"
	"github.com/pivotal-cf/go-pivnet/v4"
	"github.com/pivotal-cf/go-pivnet/v4/download"
	"github.com/pivotal-cf/go-pivnet/v4/logger"
	"io"
	"net/http"
)

const RETRY_ATTEMPTS = 3

type Client struct {
	client pivnet.Client
	logger logger.Logger
}

func NewClient(token pivnet.AccessTokenService, config pivnet.ClientConfig, logger logger.Logger) *Client {
	return &Client{
		client: pivnet.NewClient(token, config, logger),
		logger: logger,
	}
}

func (c Client) GetFederationToken(productSlug string) (pivnet.FederationToken, error) {
	var value pivnet.FederationToken
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		value, err = c.client.FederationToken.GenerateFederationToken(productSlug)
		if err == nil {
			return value, err
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}
	return value, err
}

func (c Client) ReleaseTypes() ([]pivnet.ReleaseType, error) {
	var value []pivnet.ReleaseType
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		value, err = c.client.ReleaseTypes.Get()
		if err == nil {
			return value, err
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}
	return value, err
}

func (c Client) S3PrefixForProductSlug(productSlug string) (string, error) {
	var value pivnet.Product
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		value, err = c.client.Products.Get(productSlug)
		if err == nil {
			return value.S3Directory.Path, err
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}

	return "", err
}

func (c Client) ReleasesForProductSlug(productSlug string) ([]pivnet.Release, error) {
	var value []pivnet.Release
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		value, err = c.client.Releases.List(productSlug)
		if err == nil {
			return value, err
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}
	return value, err
}

func (c Client) GetRelease(productSlug string, version string) (pivnet.Release, error) {
	var value []pivnet.Release
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		value, err = c.client.Releases.List(productSlug)
		if err == nil {
			break
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}

	if err != nil {
		return pivnet.Release{}, err
	}

	var foundRelease pivnet.Release
	for _, r := range value {
		if r.Version == version {
			foundRelease = r
			break
		}
	}

	if foundRelease.Version != version {
		return pivnet.Release{}, fmt.Errorf("release not found")
	}


	var releaseValue pivnet.Release
	for i := 0; i < RETRY_ATTEMPTS; i++ {
		releaseValue, err = c.client.Releases.Get(productSlug, foundRelease.ID)
		if err == nil {
			break
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}

	if err != nil {
		return pivnet.Release{}, err
	}
	return releaseValue, nil
}

func (c Client) UpdateRelease(productSlug string, release pivnet.Release) (pivnet.Release, error) {
	var value pivnet.Release
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		value, err = c.client.Releases.Update(productSlug, release)
		if err == nil {
			break
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}


	return value, err
}

func (c Client) CreateRelease(config pivnet.CreateReleaseConfig) (pivnet.Release, error) {
	var value pivnet.Release
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		value, err = c.client.Releases.Create(config)
		if err == nil {
			break
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}


	return value, err
}

func (c Client) DeleteRelease(productSlug string, release pivnet.Release) error {
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		err = c.client.Releases.Delete(productSlug, release)
		if err == nil {
			break
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}


	return err
}

func (c Client) AddUserGroup(productSlug string, releaseID int, userGroupID int) error {
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		err = c.client.UserGroups.AddToRelease(productSlug, releaseID, userGroupID)
		if err == nil {
			break
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}

	return err
}

func (c Client) UserGroups(productSlug string, releaseID int) ([]pivnet.UserGroup, error) {
	var value []pivnet.UserGroup
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		value, err = c.client.UserGroups.ListForRelease(productSlug, releaseID)
		if err == nil {
			break
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}


	return value, err
}

func (c Client) AcceptEULA(productSlug string, releaseID int) error {
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		err = c.client.EULA.Accept(productSlug, releaseID)
		if err == nil {
			break
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}

	return err
}

func (c Client) EULAs() ([]pivnet.EULA, error) {
	var value []pivnet.EULA
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		value, err = c.client.EULA.List()
		if err == nil {
			break
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}


	return value, err
}

func (c Client) FindProductForSlug(slug string) (pivnet.Product, error) {
	var value pivnet.Product
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		value, err = c.client.Products.Get(slug)
		if err == nil {
			break
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}


	return value, err
}

func (c Client) ProductFilesForRelease(productSlug string, releaseID int) ([]pivnet.ProductFile, error) {
	var value []pivnet.ProductFile
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		value, err = c.client.ProductFiles.ListForRelease(productSlug, releaseID)
		if err == nil {
			break
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}


	return value, err
}

func (c Client) ProductFiles(productSlug string) ([]pivnet.ProductFile, error) {
	var value []pivnet.ProductFile
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		value, err = c.client.ProductFiles.List(productSlug)
		if err == nil {
			break
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}


	return value, err
}

func (c Client) ProductFile(productSlug string, productFileID int) (pivnet.ProductFile, error) {
	var value pivnet.ProductFile
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		value, err = c.client.ProductFiles.Get(productSlug, productFileID)
		if err == nil {
			break
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}


	return value, err
}

func (c Client) ProductFileForRelease(productSlug string, releaseID int, productFileID int) (pivnet.ProductFile, error) {
	var value pivnet.ProductFile
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		value, err = c.client.ProductFiles.GetForRelease(productSlug, releaseID, productFileID)
		if err == nil {
			break
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}


	return value, err
}

func (c Client) DeleteProductFile(productSlug string, releaseID int) (pivnet.ProductFile, error) {
	var value pivnet.ProductFile
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		value, err = c.client.ProductFiles.Delete(productSlug, releaseID)
		if err == nil {
			break
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}


	return value, err
}

func (c Client) CreateProductFile(config pivnet.CreateProductFileConfig) (pivnet.ProductFile, error) {
	var value pivnet.ProductFile
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		value, err = c.client.ProductFiles.Create(config)
		if err == nil {
			break
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}


	return value, err
}

func (c Client) AddProductFile(productSlug string, releaseID int, productFileID int) error {
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		err = c.client.ProductFiles.AddToRelease(productSlug, releaseID, productFileID)
		if err == nil {
			break
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}

	return err
}

func (c Client) CreateFileGroup(config pivnet.CreateFileGroupConfig) (pivnet.FileGroup, error) {
	var value pivnet.FileGroup
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		value, err = c.client.FileGroups.Create(config)
		if err == nil {
			break
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}


	return value, err
}

func (c Client) AddToFileGroup(productSlug string, fileGroupID int, productFileID int) error {
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		err = c.client.ProductFiles.AddToFileGroup(productSlug, fileGroupID, productFileID)
		if err == nil {
			break
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}

	return err
}

func (c Client) AddFileGroup(productSlug string, releaseID int, fileGroupID int) error {
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		err = c.client.FileGroups.AddToRelease(productSlug, releaseID, fileGroupID)
		if err == nil {
			break
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}

	return err
}

func (c Client) DownloadProductFile(writer *download.FileInfo, productSlug string, releaseID int, productFileID int, progressWriter io.Writer) error {
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		err = c.client.ProductFiles.DownloadForRelease(writer, productSlug, releaseID, productFileID, progressWriter)
		if err == nil {
			break
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}

	return err
}

func (c Client) FileGroupsForRelease(productSlug string, releaseID int) ([]pivnet.FileGroup, error) {
	var value []pivnet.FileGroup
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		value, err = c.client.FileGroups.ListForRelease(productSlug, releaseID)
		if err == nil {
			break
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}


	return value, err
}

func (c Client) ImageReferences(productSlug string) ([]pivnet.ImageReference, error) {
	var value []pivnet.ImageReference
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		value, err = c.client.ImageReferences.List(productSlug)
		if err == nil {
			break
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}


	return value, err
}

func (c Client) ImageReferencesForRelease(productSlug string, releaseID int) ([]pivnet.ImageReference, error) {
	var value []pivnet.ImageReference
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		value, err = c.client.ImageReferences.ListForRelease(productSlug, releaseID)
		if err == nil {
			break
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}

	return value, err
}

func (c Client) CreateImageReference(config pivnet.CreateImageReferenceConfig) (pivnet.ImageReference, error) {
	var value pivnet.ImageReference
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		value, err = c.client.ImageReferences.Create(config)
		if err == nil {
			break
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}

	return value, err
}

func (c Client) GetImageReference(productSlug string, imageReferenceID int) (pivnet.ImageReference, error) {
	var value pivnet.ImageReference
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		value, err = c.client.ImageReferences.Get(productSlug, imageReferenceID)
		if err == nil {
			break
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}

	return value, err
}

func (c Client) AddImageReference(productSlug string, releaseID int, imageReferenceID int) error {
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		err = c.client.ImageReferences.AddToRelease(productSlug, releaseID, imageReferenceID)
		if err == nil {
			break
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}

	return err
}

func (c Client) DeleteImageReference(productSlug string, imageReferenceID int) (pivnet.ImageReference, error) {
	var value pivnet.ImageReference
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		value, err = c.client.ImageReferences.Delete(productSlug, imageReferenceID)
		if err == nil {
			break
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}

	return value, err
}

func (c Client) HelmChartReferences(productSlug string) ([]pivnet.HelmChartReference, error) {
	var value []pivnet.HelmChartReference
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		value, err = c.client.HelmChartReferences.List(productSlug)
		if err == nil {
			break
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}

	return value, err
}

func (c Client) HelmChartReferencesForRelease(productSlug string, releaseID int) ([]pivnet.HelmChartReference, error) {
	var value []pivnet.HelmChartReference
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		value, err = c.client.HelmChartReferences.ListForRelease(productSlug, releaseID)
		if err == nil {
			break
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}

	return value, err
}

func (c Client) CreateHelmChartReference(config pivnet.CreateHelmChartReferenceConfig) (pivnet.HelmChartReference, error) {
	var value pivnet.HelmChartReference
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		value, err = c.client.HelmChartReferences.Create(config)
		if err == nil {
			break
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}

	return value, err
}

func (c Client) GetHelmChartReference(productSlug string, helmChartReferenceID int) (pivnet.HelmChartReference, error) {
	var value pivnet.HelmChartReference
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		value, err = c.client.HelmChartReferences.Get(productSlug, helmChartReferenceID)
		if err == nil {
			break
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}

	return value, err
}

func (c Client) AddHelmChartReference(productSlug string, releaseID int, helmChartReferenceID int) error {
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		err = c.client.HelmChartReferences.AddToRelease(productSlug, releaseID, helmChartReferenceID)
		if err == nil {
			break
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}

	return err
}

func (c Client) DeleteHelmChartReference(productSlug string, helmChartReferenceID int) (pivnet.HelmChartReference, error) {
	var err error
	var value pivnet.HelmChartReference

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		value, err = c.client.HelmChartReferences.Delete(productSlug, helmChartReferenceID)
		if err == nil {
			break
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}

	return value, err
}

func (c Client) ReleaseDependencies(productSlug string, releaseID int) ([]pivnet.ReleaseDependency, error) {
	var value []pivnet.ReleaseDependency
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		value, err = c.client.ReleaseDependencies.List(productSlug, releaseID)
		if err == nil {
			break
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}

	return value, err
}

func (c Client) AddReleaseDependency(productSlug string, releaseID int, dependentReleaseID int) error {
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		err = c.client.ReleaseDependencies.Add(productSlug, releaseID, dependentReleaseID)
		if err == nil {
			break
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}

	return err
}

func (c Client) DependencySpecifiers(productSlug string, releaseID int) ([]pivnet.DependencySpecifier, error) {
	var value []pivnet.DependencySpecifier
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		value, err = c.client.DependencySpecifiers.List(productSlug, releaseID)
		if err == nil {
			break
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}

	return value, err
}

func (c Client) CreateDependencySpecifier(productSlug string, releaseID int, dependentProductSlug string, specifier string) (pivnet.DependencySpecifier, error) {
	var value pivnet.DependencySpecifier
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		value, err = c.client.DependencySpecifiers.Create(productSlug, releaseID, dependentProductSlug, specifier)
		if err == nil {
			break
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}

	return value, err
}

func (c Client) ReleaseUpgradePaths(productSlug string, releaseID int) ([]pivnet.ReleaseUpgradePath, error) {
	var value []pivnet.ReleaseUpgradePath
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		value, err = c.client.ReleaseUpgradePaths.Get(productSlug, releaseID)
		if err == nil {
			break
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}

	return value, err
}

func (c Client) UpgradePathSpecifiers(productSlug string, releaseID int) ([]pivnet.UpgradePathSpecifier, error) {
	var value []pivnet.UpgradePathSpecifier
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		value, err = c.client.UpgradePathSpecifiers.List(productSlug, releaseID)
		if err == nil {
			break
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}

	return value, err
}

func (c Client) CreateUpgradePathSpecifier(productSlug string, releaseID int, specifier string) (pivnet.UpgradePathSpecifier, error) {
	var value pivnet.UpgradePathSpecifier
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		value, err = c.client.UpgradePathSpecifiers.Create(productSlug, releaseID, specifier)
		if err == nil {
			break
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}

	return value, err
}

func (c Client) AddReleaseUpgradePath(productSlug string, releaseID int, previousReleaseID int) error {
	var err error

	for i := 0; i < RETRY_ATTEMPTS; i++ {
		err = c.client.ReleaseUpgradePaths.Add(productSlug, releaseID, previousReleaseID)
		if err == nil {
			break
		}
		c.logger.Info(fmt.Sprintf("Received error %s, retry attempt #%d", err.Error(), i))
	}

	return err
}

func (c Client) CreateRequest(method string, url string, body io.Reader) (*http.Request, error) {
	return c.client.CreateRequest(method, url, body)
}
