package main

import (
	"github.com/amazon-sagemaker-geospatial-capabilities/mcp-server/config"
	"github.com/amazon-sagemaker-geospatial-capabilities/mcp-server/models"
	tools_earth_observation_jobs "github.com/amazon-sagemaker-geospatial-capabilities/mcp-server/tools/earth_observation_jobs"
	tools_tile "github.com/amazon-sagemaker-geospatial-capabilities/mcp-server/tools/tile"
	tools_tags "github.com/amazon-sagemaker-geospatial-capabilities/mcp-server/tools/tags"
	tools_raster_data_collection "github.com/amazon-sagemaker-geospatial-capabilities/mcp-server/tools/raster_data_collection"
	tools_search_raster_data_collection "github.com/amazon-sagemaker-geospatial-capabilities/mcp-server/tools/search_raster_data_collection"
	tools_vector_enrichment_jobs "github.com/amazon-sagemaker-geospatial-capabilities/mcp-server/tools/vector_enrichment_jobs"
	tools_raster_data_collections "github.com/amazon-sagemaker-geospatial-capabilities/mcp-server/tools/raster_data_collections"
	tools_list_vector_enrichment_jobs "github.com/amazon-sagemaker-geospatial-capabilities/mcp-server/tools/list_vector_enrichment_jobs"
	tools_export_vector_enrichment_jobs "github.com/amazon-sagemaker-geospatial-capabilities/mcp-server/tools/export_vector_enrichment_jobs"
	tools_export_earth_observation_job "github.com/amazon-sagemaker-geospatial-capabilities/mcp-server/tools/export_earth_observation_job"
	tools_list_earth_observation_jobs "github.com/amazon-sagemaker-geospatial-capabilities/mcp-server/tools/list_earth_observation_jobs"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_earth_observation_jobs.CreateGetearthobservationjobTool(cfg),
		tools_earth_observation_jobs.CreateDeleteearthobservationjobTool(cfg),
		tools_tile.CreateGettileTool(cfg),
		tools_tags.CreateUntagresourceTool(cfg),
		tools_earth_observation_jobs.CreateStopearthobservationjobTool(cfg),
		tools_tags.CreateListtagsforresourceTool(cfg),
		tools_tags.CreateTagresourceTool(cfg),
		tools_raster_data_collection.CreateGetrasterdatacollectionTool(cfg),
		tools_search_raster_data_collection.CreateSearchrasterdatacollectionTool(cfg),
		tools_vector_enrichment_jobs.CreateStartvectorenrichmentjobTool(cfg),
		tools_earth_observation_jobs.CreateStartearthobservationjobTool(cfg),
		tools_raster_data_collections.CreateListrasterdatacollectionsTool(cfg),
		tools_vector_enrichment_jobs.CreateDeletevectorenrichmentjobTool(cfg),
		tools_vector_enrichment_jobs.CreateGetvectorenrichmentjobTool(cfg),
		tools_list_vector_enrichment_jobs.CreateListvectorenrichmentjobsTool(cfg),
		tools_vector_enrichment_jobs.CreateStopvectorenrichmentjobTool(cfg),
		tools_export_vector_enrichment_jobs.CreateExportvectorenrichmentjobTool(cfg),
		tools_export_earth_observation_job.CreateExportearthobservationjobTool(cfg),
		tools_list_earth_observation_jobs.CreateListearthobservationjobsTool(cfg),
	}
}
