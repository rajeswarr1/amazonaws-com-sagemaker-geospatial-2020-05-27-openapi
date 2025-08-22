package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// GetRasterDataCollectionInput represents the GetRasterDataCollectionInput schema from the OpenAPI specification
type GetRasterDataCollectionInput struct {
}

// GetEarthObservationJobOutput represents the GetEarthObservationJobOutput schema from the OpenAPI specification
type GetEarthObservationJobOutput struct {
	Errordetails interface{} `json:"ErrorDetails,omitempty"`
	Inputconfig interface{} `json:"InputConfig"`
	Name interface{} `json:"Name"`
	Tags interface{} `json:"Tags,omitempty"`
	Arn interface{} `json:"Arn"`
	Executionrolearn interface{} `json:"ExecutionRoleArn,omitempty"`
	Exportstatus interface{} `json:"ExportStatus,omitempty"`
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
	Creationtime interface{} `json:"CreationTime"`
	Exporterrordetails interface{} `json:"ExportErrorDetails,omitempty"`
	Durationinseconds interface{} `json:"DurationInSeconds"`
	Jobconfig interface{} `json:"JobConfig"`
	Outputbands interface{} `json:"OutputBands,omitempty"`
	Status interface{} `json:"Status"`
}

// ListVectorEnrichmentJobOutput represents the ListVectorEnrichmentJobOutput schema from the OpenAPI specification
type ListVectorEnrichmentJobOutput struct {
	Vectorenrichmentjobsummaries interface{} `json:"VectorEnrichmentJobSummaries"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ViewSunAzimuthInput represents the ViewSunAzimuthInput schema from the OpenAPI specification
type ViewSunAzimuthInput struct {
	Lowerbound interface{} `json:"LowerBound"`
	Upperbound interface{} `json:"UpperBound"`
}

// GetTileOutput represents the GetTileOutput schema from the OpenAPI specification
type GetTileOutput struct {
	Binaryfile interface{} `json:"BinaryFile,omitempty"`
}

// ExportErrorDetailsOutput represents the ExportErrorDetailsOutput schema from the OpenAPI specification
type ExportErrorDetailsOutput struct {
	Message interface{} `json:"Message,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// ExportVectorEnrichmentJobOutput represents the ExportVectorEnrichmentJobOutput schema from the OpenAPI specification
type ExportVectorEnrichmentJobOutput struct {
	Arn interface{} `json:"Arn"`
	Creationtime interface{} `json:"CreationTime"`
	Executionrolearn interface{} `json:"ExecutionRoleArn"`
	Exportstatus interface{} `json:"ExportStatus"`
	Outputconfig interface{} `json:"OutputConfig"`
}

// CustomIndicesInput represents the CustomIndicesInput schema from the OpenAPI specification
type CustomIndicesInput struct {
	Operations interface{} `json:"Operations,omitempty"`
}

// RasterDataCollectionQueryInput represents the RasterDataCollectionQueryInput schema from the OpenAPI specification
type RasterDataCollectionQueryInput struct {
	Timerangefilter interface{} `json:"TimeRangeFilter"`
	Areaofinterest interface{} `json:"AreaOfInterest,omitempty"`
	Propertyfilters interface{} `json:"PropertyFilters,omitempty"`
	Rasterdatacollectionarn interface{} `json:"RasterDataCollectionArn"`
}

// ResamplingConfigInput represents the ResamplingConfigInput schema from the OpenAPI specification
type ResamplingConfigInput struct {
	Targetbands interface{} `json:"TargetBands,omitempty"`
	Algorithmname interface{} `json:"AlgorithmName,omitempty"`
	Outputresolution interface{} `json:"OutputResolution"`
}

// StopVectorEnrichmentJobOutput represents the StopVectorEnrichmentJobOutput schema from the OpenAPI specification
type StopVectorEnrichmentJobOutput struct {
}

// ListVectorEnrichmentJobOutputConfig represents the ListVectorEnrichmentJobOutputConfig schema from the OpenAPI specification
type ListVectorEnrichmentJobOutputConfig struct {
	Tags interface{} `json:"Tags,omitempty"`
	TypeField interface{} `json:"Type"`
	Arn interface{} `json:"Arn"`
	Creationtime interface{} `json:"CreationTime"`
	Durationinseconds interface{} `json:"DurationInSeconds"`
	Name interface{} `json:"Name"`
	Status interface{} `json:"Status"`
}

// CloudRemovalConfigInput represents the CloudRemovalConfigInput schema from the OpenAPI specification
type CloudRemovalConfigInput struct {
	Algorithmname interface{} `json:"AlgorithmName,omitempty"`
	Interpolationvalue interface{} `json:"InterpolationValue,omitempty"`
	Targetbands interface{} `json:"TargetBands,omitempty"`
}

// StopEarthObservationJobInput represents the StopEarthObservationJobInput schema from the OpenAPI specification
type StopEarthObservationJobInput struct {
	Arn interface{} `json:"Arn"`
}

// PropertyFilters represents the PropertyFilters schema from the OpenAPI specification
type PropertyFilters struct {
	Logicaloperator interface{} `json:"LogicalOperator,omitempty"`
	Properties interface{} `json:"Properties,omitempty"`
}

// TagResourceResponse represents the TagResourceResponse schema from the OpenAPI specification
type TagResourceResponse struct {
}

// RasterDataCollectionQueryOutput represents the RasterDataCollectionQueryOutput schema from the OpenAPI specification
type RasterDataCollectionQueryOutput struct {
	Timerangefilter interface{} `json:"TimeRangeFilter"`
	Areaofinterest interface{} `json:"AreaOfInterest,omitempty"`
	Propertyfilters interface{} `json:"PropertyFilters,omitempty"`
	Rasterdatacollectionarn interface{} `json:"RasterDataCollectionArn"`
	Rasterdatacollectionname interface{} `json:"RasterDataCollectionName"`
}

// RasterDataCollectionQueryWithBandFilterInput represents the RasterDataCollectionQueryWithBandFilterInput schema from the OpenAPI specification
type RasterDataCollectionQueryWithBandFilterInput struct {
	Areaofinterest interface{} `json:"AreaOfInterest,omitempty"`
	Bandfilter interface{} `json:"BandFilter,omitempty"`
	Propertyfilters interface{} `json:"PropertyFilters,omitempty"`
	Timerangefilter interface{} `json:"TimeRangeFilter"`
}

// ListEarthObservationJobOutputConfig represents the ListEarthObservationJobOutputConfig schema from the OpenAPI specification
type ListEarthObservationJobOutputConfig struct {
	Arn interface{} `json:"Arn"`
	Creationtime interface{} `json:"CreationTime"`
	Durationinseconds interface{} `json:"DurationInSeconds"`
	Name interface{} `json:"Name"`
	Operationtype interface{} `json:"OperationType"`
	Status interface{} `json:"Status"`
	Tags interface{} `json:"Tags,omitempty"`
}

// UntagResourceResponse represents the UntagResourceResponse schema from the OpenAPI specification
type UntagResourceResponse struct {
}

// SearchRasterDataCollectionOutput represents the SearchRasterDataCollectionOutput schema from the OpenAPI specification
type SearchRasterDataCollectionOutput struct {
	Approximateresultcount interface{} `json:"ApproximateResultCount"`
	Items interface{} `json:"Items,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// GetRasterDataCollectionOutput represents the GetRasterDataCollectionOutput schema from the OpenAPI specification
type GetRasterDataCollectionOutput struct {
	Imagesourcebands interface{} `json:"ImageSourceBands"`
	Name interface{} `json:"Name"`
	Supportedfilters interface{} `json:"SupportedFilters"`
	Tags interface{} `json:"Tags,omitempty"`
	TypeField interface{} `json:"Type"`
	Arn interface{} `json:"Arn"`
	Description interface{} `json:"Description"`
	Descriptionpageurl interface{} `json:"DescriptionPageUrl"`
}

// TagResourceRequest represents the TagResourceRequest schema from the OpenAPI specification
type TagResourceRequest struct {
	Tags interface{} `json:"Tags"`
}

// InputConfigInput represents the InputConfigInput schema from the OpenAPI specification
type InputConfigInput struct {
	Previousearthobservationjobarn interface{} `json:"PreviousEarthObservationJobArn,omitempty"`
	Rasterdatacollectionquery interface{} `json:"RasterDataCollectionQuery,omitempty"`
}

// PolygonGeometryInput represents the PolygonGeometryInput schema from the OpenAPI specification
type PolygonGeometryInput struct {
	Coordinates interface{} `json:"Coordinates"`
}

// UntagResourceRequest represents the UntagResourceRequest schema from the OpenAPI specification
type UntagResourceRequest struct {
}

// Properties represents the Properties schema from the OpenAPI specification
type Properties struct {
	Landsatcloudcoverland interface{} `json:"LandsatCloudCoverLand,omitempty"`
	Platform interface{} `json:"Platform,omitempty"`
	Viewoffnadir interface{} `json:"ViewOffNadir,omitempty"`
	Viewsunazimuth interface{} `json:"ViewSunAzimuth,omitempty"`
	Viewsunelevation interface{} `json:"ViewSunElevation,omitempty"`
	Eocloudcover interface{} `json:"EoCloudCover,omitempty"`
}

// VectorEnrichmentJobDataSourceConfigInput represents the VectorEnrichmentJobDataSourceConfigInput schema from the OpenAPI specification
type VectorEnrichmentJobDataSourceConfigInput struct {
	S3data interface{} `json:"S3Data,omitempty"`
}

// AssetValue represents the AssetValue schema from the OpenAPI specification
type AssetValue struct {
	Href interface{} `json:"Href,omitempty"`
}

// GetVectorEnrichmentJobOutput represents the GetVectorEnrichmentJobOutput schema from the OpenAPI specification
type GetVectorEnrichmentJobOutput struct {
	Tags interface{} `json:"Tags,omitempty"`
	Durationinseconds interface{} `json:"DurationInSeconds"`
	Exporterrordetails interface{} `json:"ExportErrorDetails,omitempty"`
	Exportstatus interface{} `json:"ExportStatus,omitempty"`
	Jobconfig interface{} `json:"JobConfig"`
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
	TypeField interface{} `json:"Type"`
	Arn interface{} `json:"Arn"`
	Inputconfig interface{} `json:"InputConfig"`
	Name interface{} `json:"Name"`
	Status interface{} `json:"Status"`
	Creationtime interface{} `json:"CreationTime"`
	Errordetails interface{} `json:"ErrorDetails,omitempty"`
	Executionrolearn interface{} `json:"ExecutionRoleArn"`
}

// SearchRasterDataCollectionInput represents the SearchRasterDataCollectionInput schema from the OpenAPI specification
type SearchRasterDataCollectionInput struct {
	Arn interface{} `json:"Arn"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Rasterdatacollectionquery interface{} `json:"RasterDataCollectionQuery"`
}

// StartEarthObservationJobOutput represents the StartEarthObservationJobOutput schema from the OpenAPI specification
type StartEarthObservationJobOutput struct {
	Jobconfig interface{} `json:"JobConfig"`
	Creationtime interface{} `json:"CreationTime"`
	Name interface{} `json:"Name"`
	Arn interface{} `json:"Arn"`
	Inputconfig interface{} `json:"InputConfig,omitempty"`
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
	Status interface{} `json:"Status"`
	Tags interface{} `json:"Tags,omitempty"`
	Durationinseconds interface{} `json:"DurationInSeconds"`
	Executionrolearn interface{} `json:"ExecutionRoleArn"`
}

// StopVectorEnrichmentJobInput represents the StopVectorEnrichmentJobInput schema from the OpenAPI specification
type StopVectorEnrichmentJobInput struct {
	Arn interface{} `json:"Arn"`
}

// ListRasterDataCollectionsInput represents the ListRasterDataCollectionsInput schema from the OpenAPI specification
type ListRasterDataCollectionsInput struct {
}

// Operation represents the Operation schema from the OpenAPI specification
type Operation struct {
	Equation interface{} `json:"Equation"`
	Name interface{} `json:"Name"`
	Outputtype interface{} `json:"OutputType,omitempty"`
}

// VectorEnrichmentJobConfig represents the VectorEnrichmentJobConfig schema from the OpenAPI specification
type VectorEnrichmentJobConfig struct {
	Reversegeocodingconfig interface{} `json:"ReverseGeocodingConfig,omitempty"`
	Mapmatchingconfig interface{} `json:"MapMatchingConfig,omitempty"`
}

// ViewOffNadirInput represents the ViewOffNadirInput schema from the OpenAPI specification
type ViewOffNadirInput struct {
	Upperbound interface{} `json:"UpperBound"`
	Lowerbound interface{} `json:"LowerBound"`
}

// VectorEnrichmentJobExportErrorDetails represents the VectorEnrichmentJobExportErrorDetails schema from the OpenAPI specification
type VectorEnrichmentJobExportErrorDetails struct {
	Message interface{} `json:"Message,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// UserDefined represents the UserDefined schema from the OpenAPI specification
type UserDefined struct {
	Unit interface{} `json:"Unit"`
	Value interface{} `json:"Value"`
}

// DeleteEarthObservationJobInput represents the DeleteEarthObservationJobInput schema from the OpenAPI specification
type DeleteEarthObservationJobInput struct {
}

// ListTagsForResourceResponse represents the ListTagsForResourceResponse schema from the OpenAPI specification
type ListTagsForResourceResponse struct {
	Tags interface{} `json:"Tags,omitempty"`
}

// LandCoverSegmentationConfigInput represents the LandCoverSegmentationConfigInput schema from the OpenAPI specification
type LandCoverSegmentationConfigInput struct {
}

// PropertyFilter represents the PropertyFilter schema from the OpenAPI specification
type PropertyFilter struct {
	Property interface{} `json:"Property"`
}

// PlatformInput represents the PlatformInput schema from the OpenAPI specification
type PlatformInput struct {
	Value interface{} `json:"Value"`
	Comparisonoperator interface{} `json:"ComparisonOperator,omitempty"`
}

// ListEarthObservationJobInput represents the ListEarthObservationJobInput schema from the OpenAPI specification
type ListEarthObservationJobInput struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Statusequals interface{} `json:"StatusEquals,omitempty"`
}

// ListVectorEnrichmentJobInput represents the ListVectorEnrichmentJobInput schema from the OpenAPI specification
type ListVectorEnrichmentJobInput struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Statusequals interface{} `json:"StatusEquals,omitempty"`
}

// VectorEnrichmentJobS3Data represents the VectorEnrichmentJobS3Data schema from the OpenAPI specification
type VectorEnrichmentJobS3Data struct {
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
	S3uri interface{} `json:"S3Uri"`
}

// ListEarthObservationJobOutput represents the ListEarthObservationJobOutput schema from the OpenAPI specification
type ListEarthObservationJobOutput struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Earthobservationjobsummaries interface{} `json:"EarthObservationJobSummaries"`
}

// TemporalStatisticsConfigInput represents the TemporalStatisticsConfigInput schema from the OpenAPI specification
type TemporalStatisticsConfigInput struct {
	Statistics interface{} `json:"Statistics"`
	Targetbands interface{} `json:"TargetBands,omitempty"`
	Groupby interface{} `json:"GroupBy,omitempty"`
}

// OutputResolutionResamplingInput represents the OutputResolutionResamplingInput schema from the OpenAPI specification
type OutputResolutionResamplingInput struct {
	Userdefined interface{} `json:"UserDefined"`
}

// EarthObservationJobErrorDetails represents the EarthObservationJobErrorDetails schema from the OpenAPI specification
type EarthObservationJobErrorDetails struct {
	Message interface{} `json:"Message,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// InputConfigOutput represents the InputConfigOutput schema from the OpenAPI specification
type InputConfigOutput struct {
	Rasterdatacollectionquery interface{} `json:"RasterDataCollectionQuery,omitempty"`
	Previousearthobservationjobarn interface{} `json:"PreviousEarthObservationJobArn,omitempty"`
}

// Tags represents the Tags schema from the OpenAPI specification
type Tags struct {
}

// OutputConfigInput represents the OutputConfigInput schema from the OpenAPI specification
type OutputConfigInput struct {
	S3data interface{} `json:"S3Data"`
}

// ItemSource represents the ItemSource schema from the OpenAPI specification
type ItemSource struct {
	Datetime interface{} `json:"DateTime"`
	Geometry interface{} `json:"Geometry"`
	Id interface{} `json:"Id"`
	Properties interface{} `json:"Properties,omitempty"`
	Assets interface{} `json:"Assets,omitempty"`
}

// ListRasterDataCollectionsOutput represents the ListRasterDataCollectionsOutput schema from the OpenAPI specification
type ListRasterDataCollectionsOutput struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Rasterdatacollectionsummaries interface{} `json:"RasterDataCollectionSummaries"`
}

// Property represents the Property schema from the OpenAPI specification
type Property struct {
	Viewsunelevation interface{} `json:"ViewSunElevation,omitempty"`
	Eocloudcover interface{} `json:"EoCloudCover,omitempty"`
	Landsatcloudcoverland interface{} `json:"LandsatCloudCoverLand,omitempty"`
	Platform interface{} `json:"Platform,omitempty"`
	Viewoffnadir interface{} `json:"ViewOffNadir,omitempty"`
	Viewsunazimuth interface{} `json:"ViewSunAzimuth,omitempty"`
}

// StartVectorEnrichmentJobInput represents the StartVectorEnrichmentJobInput schema from the OpenAPI specification
type StartVectorEnrichmentJobInput struct {
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
	Name interface{} `json:"Name"`
	Tags interface{} `json:"Tags,omitempty"`
	Clienttoken interface{} `json:"ClientToken,omitempty"`
	Executionrolearn interface{} `json:"ExecutionRoleArn"`
	Inputconfig interface{} `json:"InputConfig"`
	Jobconfig interface{} `json:"JobConfig"`
}

// GetEarthObservationJobInput represents the GetEarthObservationJobInput schema from the OpenAPI specification
type GetEarthObservationJobInput struct {
}

// ExportVectorEnrichmentJobInput represents the ExportVectorEnrichmentJobInput schema from the OpenAPI specification
type ExportVectorEnrichmentJobInput struct {
	Arn interface{} `json:"Arn"`
	Clienttoken interface{} `json:"ClientToken,omitempty"`
	Executionrolearn interface{} `json:"ExecutionRoleArn"`
	Outputconfig interface{} `json:"OutputConfig"`
}

// TimeRangeFilterInput represents the TimeRangeFilterInput schema from the OpenAPI specification
type TimeRangeFilterInput struct {
	Endtime interface{} `json:"EndTime"`
	Starttime interface{} `json:"StartTime"`
}

// ExportEarthObservationJobInput represents the ExportEarthObservationJobInput schema from the OpenAPI specification
type ExportEarthObservationJobInput struct {
	Outputconfig interface{} `json:"OutputConfig"`
	Arn interface{} `json:"Arn"`
	Clienttoken interface{} `json:"ClientToken,omitempty"`
	Executionrolearn interface{} `json:"ExecutionRoleArn"`
	Exportsourceimages interface{} `json:"ExportSourceImages,omitempty"`
}

// OutputBand represents the OutputBand schema from the OpenAPI specification
type OutputBand struct {
	Outputdatatype interface{} `json:"OutputDataType"`
	Bandname interface{} `json:"BandName"`
}

// GetTileInput represents the GetTileInput schema from the OpenAPI specification
type GetTileInput struct {
}

// AreaOfInterest represents the AreaOfInterest schema from the OpenAPI specification
type AreaOfInterest struct {
	Areaofinterestgeometry interface{} `json:"AreaOfInterestGeometry,omitempty"`
}

// ExportVectorEnrichmentJobOutputConfig represents the ExportVectorEnrichmentJobOutputConfig schema from the OpenAPI specification
type ExportVectorEnrichmentJobOutputConfig struct {
	S3data interface{} `json:"S3Data"`
}

// ListTagsForResourceRequest represents the ListTagsForResourceRequest schema from the OpenAPI specification
type ListTagsForResourceRequest struct {
}

// CloudMaskingConfigInput represents the CloudMaskingConfigInput schema from the OpenAPI specification
type CloudMaskingConfigInput struct {
}

// StackConfigInput represents the StackConfigInput schema from the OpenAPI specification
type StackConfigInput struct {
	Outputresolution interface{} `json:"OutputResolution,omitempty"`
	Targetbands interface{} `json:"TargetBands,omitempty"`
}

// AreaOfInterestGeometry represents the AreaOfInterestGeometry schema from the OpenAPI specification
type AreaOfInterestGeometry struct {
	Polygongeometry interface{} `json:"PolygonGeometry,omitempty"`
	Multipolygongeometry interface{} `json:"MultiPolygonGeometry,omitempty"`
}

// DeleteEarthObservationJobOutput represents the DeleteEarthObservationJobOutput schema from the OpenAPI specification
type DeleteEarthObservationJobOutput struct {
}

// ReverseGeocodingConfig represents the ReverseGeocodingConfig schema from the OpenAPI specification
type ReverseGeocodingConfig struct {
	Xattributename interface{} `json:"XAttributeName"`
	Yattributename interface{} `json:"YAttributeName"`
}

// StartVectorEnrichmentJobOutput represents the StartVectorEnrichmentJobOutput schema from the OpenAPI specification
type StartVectorEnrichmentJobOutput struct {
	Executionrolearn interface{} `json:"ExecutionRoleArn"`
	Inputconfig interface{} `json:"InputConfig"`
	Jobconfig interface{} `json:"JobConfig"`
	Tags interface{} `json:"Tags,omitempty"`
	Status interface{} `json:"Status"`
	Durationinseconds interface{} `json:"DurationInSeconds"`
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
	Name interface{} `json:"Name"`
	Creationtime interface{} `json:"CreationTime"`
	TypeField interface{} `json:"Type"`
	Arn interface{} `json:"Arn"`
}

// GeoMosaicConfigInput represents the GeoMosaicConfigInput schema from the OpenAPI specification
type GeoMosaicConfigInput struct {
	Algorithmname interface{} `json:"AlgorithmName,omitempty"`
	Targetbands interface{} `json:"TargetBands,omitempty"`
}

// ViewSunElevationInput represents the ViewSunElevationInput schema from the OpenAPI specification
type ViewSunElevationInput struct {
	Lowerbound interface{} `json:"LowerBound"`
	Upperbound interface{} `json:"UpperBound"`
}

// ZonalStatisticsConfigInput represents the ZonalStatisticsConfigInput schema from the OpenAPI specification
type ZonalStatisticsConfigInput struct {
	Zones3pathkmskeyid interface{} `json:"ZoneS3PathKmsKeyId,omitempty"`
	Statistics interface{} `json:"Statistics"`
	Targetbands interface{} `json:"TargetBands,omitempty"`
	Zones3path interface{} `json:"ZoneS3Path"`
}

// MapMatchingConfig represents the MapMatchingConfig schema from the OpenAPI specification
type MapMatchingConfig struct {
	Timestampattributename interface{} `json:"TimestampAttributeName"`
	Xattributename interface{} `json:"XAttributeName"`
	Yattributename interface{} `json:"YAttributeName"`
	Idattributename interface{} `json:"IdAttributeName"`
}

// AssetsMap represents the AssetsMap schema from the OpenAPI specification
type AssetsMap struct {
}

// DeleteVectorEnrichmentJobOutput represents the DeleteVectorEnrichmentJobOutput schema from the OpenAPI specification
type DeleteVectorEnrichmentJobOutput struct {
}

// Geometry represents the Geometry schema from the OpenAPI specification
type Geometry struct {
	Coordinates interface{} `json:"Coordinates"`
	TypeField interface{} `json:"Type"`
}

// ExportErrorDetails represents the ExportErrorDetails schema from the OpenAPI specification
type ExportErrorDetails struct {
	Exportresults interface{} `json:"ExportResults,omitempty"`
	Exportsourceimages interface{} `json:"ExportSourceImages,omitempty"`
}

// BandMathConfigInput represents the BandMathConfigInput schema from the OpenAPI specification
type BandMathConfigInput struct {
	Customindices interface{} `json:"CustomIndices,omitempty"`
	Predefinedindices interface{} `json:"PredefinedIndices,omitempty"`
}

// EoCloudCoverInput represents the EoCloudCoverInput schema from the OpenAPI specification
type EoCloudCoverInput struct {
	Lowerbound interface{} `json:"LowerBound"`
	Upperbound interface{} `json:"UpperBound"`
}

// VectorEnrichmentJobInputConfig represents the VectorEnrichmentJobInputConfig schema from the OpenAPI specification
type VectorEnrichmentJobInputConfig struct {
	Datasourceconfig interface{} `json:"DataSourceConfig"`
	Documenttype interface{} `json:"DocumentType"`
}

// JobConfigInput represents the JobConfigInput schema from the OpenAPI specification
type JobConfigInput struct {
	Temporalstatisticsconfig interface{} `json:"TemporalStatisticsConfig,omitempty"`
	Zonalstatisticsconfig interface{} `json:"ZonalStatisticsConfig,omitempty"`
	Cloudmaskingconfig interface{} `json:"CloudMaskingConfig,omitempty"`
	Cloudremovalconfig interface{} `json:"CloudRemovalConfig,omitempty"`
	Stackconfig interface{} `json:"StackConfig,omitempty"`
	Geomosaicconfig interface{} `json:"GeoMosaicConfig,omitempty"`
	Resamplingconfig interface{} `json:"ResamplingConfig,omitempty"`
	Bandmathconfig interface{} `json:"BandMathConfig,omitempty"`
	Landcoversegmentationconfig interface{} `json:"LandCoverSegmentationConfig,omitempty"`
}

// DeleteVectorEnrichmentJobInput represents the DeleteVectorEnrichmentJobInput schema from the OpenAPI specification
type DeleteVectorEnrichmentJobInput struct {
}

// MultiPolygonGeometryInput represents the MultiPolygonGeometryInput schema from the OpenAPI specification
type MultiPolygonGeometryInput struct {
	Coordinates interface{} `json:"Coordinates"`
}

// ExportS3DataInput represents the ExportS3DataInput schema from the OpenAPI specification
type ExportS3DataInput struct {
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
	S3uri interface{} `json:"S3Uri"`
}

// OutputResolutionStackInput represents the OutputResolutionStackInput schema from the OpenAPI specification
type OutputResolutionStackInput struct {
	Predefined interface{} `json:"Predefined,omitempty"`
	Userdefined interface{} `json:"UserDefined,omitempty"`
}

// Filter represents the Filter schema from the OpenAPI specification
type Filter struct {
	Maximum interface{} `json:"Maximum,omitempty"`
	Minimum interface{} `json:"Minimum,omitempty"`
	Name interface{} `json:"Name"`
	TypeField interface{} `json:"Type"`
}

// TimeRangeFilterOutput represents the TimeRangeFilterOutput schema from the OpenAPI specification
type TimeRangeFilterOutput struct {
	Endtime interface{} `json:"EndTime"`
	Starttime interface{} `json:"StartTime"`
}

// GetVectorEnrichmentJobInput represents the GetVectorEnrichmentJobInput schema from the OpenAPI specification
type GetVectorEnrichmentJobInput struct {
}

// VectorEnrichmentJobErrorDetails represents the VectorEnrichmentJobErrorDetails schema from the OpenAPI specification
type VectorEnrichmentJobErrorDetails struct {
	Errormessage interface{} `json:"ErrorMessage,omitempty"`
	Errortype interface{} `json:"ErrorType,omitempty"`
}

// LandsatCloudCoverLandInput represents the LandsatCloudCoverLandInput schema from the OpenAPI specification
type LandsatCloudCoverLandInput struct {
	Lowerbound interface{} `json:"LowerBound"`
	Upperbound interface{} `json:"UpperBound"`
}

// RasterDataCollectionMetadata represents the RasterDataCollectionMetadata schema from the OpenAPI specification
type RasterDataCollectionMetadata struct {
	Tags interface{} `json:"Tags,omitempty"`
	TypeField interface{} `json:"Type"`
	Arn interface{} `json:"Arn"`
	Description interface{} `json:"Description"`
	Descriptionpageurl interface{} `json:"DescriptionPageUrl,omitempty"`
	Name interface{} `json:"Name"`
	Supportedfilters interface{} `json:"SupportedFilters"`
}

// StartEarthObservationJobInput represents the StartEarthObservationJobInput schema from the OpenAPI specification
type StartEarthObservationJobInput struct {
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
	Name interface{} `json:"Name"`
	Tags interface{} `json:"Tags,omitempty"`
	Clienttoken interface{} `json:"ClientToken,omitempty"`
	Executionrolearn interface{} `json:"ExecutionRoleArn"`
	Inputconfig interface{} `json:"InputConfig"`
	Jobconfig interface{} `json:"JobConfig"`
}

// ExportEarthObservationJobOutput represents the ExportEarthObservationJobOutput schema from the OpenAPI specification
type ExportEarthObservationJobOutput struct {
	Outputconfig interface{} `json:"OutputConfig"`
	Arn interface{} `json:"Arn"`
	Creationtime interface{} `json:"CreationTime"`
	Executionrolearn interface{} `json:"ExecutionRoleArn"`
	Exportsourceimages interface{} `json:"ExportSourceImages,omitempty"`
	Exportstatus interface{} `json:"ExportStatus"`
}

// StopEarthObservationJobOutput represents the StopEarthObservationJobOutput schema from the OpenAPI specification
type StopEarthObservationJobOutput struct {
}
