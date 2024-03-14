package main

import (
	"context"
	"fmt"

	"github.com/hasura/ndc-sdk-go/connector"
	"github.com/hasura/ndc-sdk-go/schema"
	"github.com/hasura/ndc-sdk-go/utils"
)

type Configuration struct{}
type State struct{}
type Connector struct{}

func (c *Connector) ParseConfiguration(configurationDir string) (*Configuration, error) {
	return &Configuration{}, nil
}

func (c *Connector) TryInitState(configuration *Configuration, metrics *connector.TelemetryState) (*State, error) {
	return &State{}, nil
}

func (c *Connector) HealthCheck(ctx context.Context, configuration *Configuration, state *State) error {
	return nil
}

func (c *Connector) GetCapabilities(configuration *Configuration) *schema.CapabilitiesResponse {
	return &schema.CapabilitiesResponse{
		Version: "0.1.0",
		Capabilities: schema.Capabilities{
			Query: schema.QueryCapabilities{
				Variables: schema.LeafCapability{},
			},
			Mutation: schema.MutationCapabilities{},
		},
	}
}

func (c *Connector) GetSchema(configuration *Configuration) (*schema.SchemaResponse, error) {
	ndcSchema := &schema.SchemaResponse{
		ScalarTypes: schema.SchemaResponseScalarTypes{
			"Int": schema.ScalarType{
				AggregateFunctions: schema.ScalarTypeAggregateFunctions{
					"max": schema.AggregateFunctionDefinition{
						ResultType: schema.NewNullableNamedType("Int").Encode(),
					},
					"min": schema.AggregateFunctionDefinition{
						ResultType: schema.NewNullableNamedType("Int").Encode(),
					},
				},
				ComparisonOperators: map[string]schema.ComparisonOperatorDefinition{
					"eq": schema.NewComparisonOperatorEqual().Encode(),
					"in": schema.NewComparisonOperatorIn().Encode(),
				},
			},
			"Float": schema.ScalarType{
				AggregateFunctions: schema.ScalarTypeAggregateFunctions{
					"max": schema.AggregateFunctionDefinition{
						ResultType: schema.NewNullableNamedType("Float").Encode(),
					},
					"min": schema.AggregateFunctionDefinition{
						ResultType: schema.NewNullableNamedType("Float").Encode(),
					},
				},
				ComparisonOperators: map[string]schema.ComparisonOperatorDefinition{
					"eq": schema.NewComparisonOperatorEqual().Encode(),
					"in": schema.NewComparisonOperatorIn().Encode(),
				},
			},
			"String": schema.ScalarType{
				AggregateFunctions: schema.ScalarTypeAggregateFunctions{},
				ComparisonOperators: map[string]schema.ComparisonOperatorDefinition{
					"eq":   schema.NewComparisonOperatorEqual().Encode(),
					"in":   schema.NewComparisonOperatorIn().Encode(),
					"like": schema.NewComparisonOperatorCustom(schema.NewNamedType("String")).Encode(),
				},
			},
		},
		ObjectTypes: schema.SchemaResponseObjectTypes{
			"kibana_sample_data_ecommerce": schema.ObjectType{
				Description: nil,
				Fields: schema.ObjectTypeFields{
					"_id":      {Type: schema.NewNamedType("String").Encode()},
					"category": {Type: schema.NewNamedType("String").Encode()},
					"currency": {Type: schema.NewNamedType("String").Encode()},
					"customer_birth_date": {
						Type:        schema.NewNamedType("String").Encode(),
						Description: utils.ToPtr("handle date object"),
					},
					"customer_first_name": {Type: schema.NewNamedType("String").Encode()},
					"customer_full_name":  {Type: schema.NewNamedType("String").Encode()},
					"customer_gender":     {Type: schema.NewNamedType("String").Encode()},
					"customer_id":         {Type: schema.NewNamedType("String").Encode()},
					"customer_last_name":  {Type: schema.NewNamedType("String").Encode()},
					"customer_phone":      {Type: schema.NewNamedType("String").Encode()},
					"day_of_week":         {Type: schema.NewNamedType("String").Encode()},
					"day_of_week_i":       {Type: schema.NewNamedType("Int").Encode()},
					"event":               {Type: schema.NewNamedType("event").Encode()},
					"geoip":               {Type: schema.NewNamedType("geoip").Encode()},
					"manufacturer":        {Type: schema.NewNamedType("String").Encode()},
					"order_date": {
						Type:        schema.NewNamedType("String").Encode(),
						Description: utils.ToPtr("handle date object"),
					},
					"products":              {Type: schema.NewNamedType("products").Encode()},
					"sku":                   {Type: schema.NewArrayType(schema.NewNamedType("String")).Encode()},
					"taxful_total_price":    {Type: schema.NewNamedType("Float").Encode()},
					"taxless_total_price":   {Type: schema.NewNamedType("Float").Encode()},
					"total_quantity":        {Type: schema.NewNamedType("Int").Encode()},
					"total_unique_products": {Type: schema.NewNamedType("Int").Encode()},
					"type":                  {Type: schema.NewNamedType("String").Encode()},
					"user":                  {Type: schema.NewNamedType("String").Encode()},
				},
			},
			"event": schema.ObjectType{
				Fields: schema.ObjectTypeFields{
					"dataset": {Type: schema.NewNamedType("String").Encode()},
				},
			},
			"geoip": schema.ObjectType{
				Fields: schema.ObjectTypeFields{
					"city_name":        {Type: schema.NewNamedType("String").Encode()},
					"continent_name":   {Type: schema.NewNamedType("String").Encode()},
					"country_iso_code": {Type: schema.NewNamedType("String").Encode()},
					"location":         {Type: schema.NewNamedType("String").Encode()},
					"region_name":      {Type: schema.NewNamedType("String").Encode()},
				},
			},
			"products": schema.ObjectType{
				Fields: schema.ObjectTypeFields{
					"_id":             {Type: schema.NewNamedType("String").Encode()},
					"base_price":      {Type: schema.NewNamedType("Float").Encode()},
					"base_unit_price": {Type: schema.NewNamedType("Float").Encode()},
					"category":        {Type: schema.NewNamedType("String").Encode()},
					"created_on": {
						Description: utils.ToPtr("handle date object"),
						Type:        schema.NewNamedType("Float").Encode(),
					},
					"discount_amount":      {Type: schema.NewNamedType("Float").Encode()},
					"discount_percentage":  {Type: schema.NewNamedType("Float").Encode()},
					"manufacturer":         {Type: schema.NewNamedType("String").Encode()},
					"min_price":            {Type: schema.NewNamedType("Float").Encode()},
					"price":                {Type: schema.NewNamedType("Float").Encode()},
					"product_name":         {Type: schema.NewNamedType("String").Encode()},
					"quantity":             {Type: schema.NewNamedType("Int").Encode()},
					"sku":                  {Type: schema.NewNamedType("String").Encode()},
					"tax_amount":           {Type: schema.NewNamedType("Float").Encode()},
					"taxful_price":         {Type: schema.NewNamedType("Float").Encode()},
					"taxless_price":        {Type: schema.NewNamedType("Float").Encode()},
					"unit_discount_amount": {Type: schema.NewNamedType("Float").Encode()},
					"product_id":           {Type: schema.NewNamedType("Int").Encode()}},
			},
		},
		Collections: []schema.CollectionInfo{
			{
				Name:        "kibana_sample_data_ecommerce",
				Description: utils.ToPtr("A collection of ecommerce data"),
				Arguments:   schema.CollectionInfoArguments{},
				Type:        "kibana_sample_data_ecommerce",
				UniquenessConstraints: schema.CollectionInfoUniquenessConstraints{
					"EcommerceByID": schema.UniquenessConstraint{
						UniqueColumns: []string{"id"},
					},
				},
				ForeignKeys: schema.CollectionInfoForeignKeys{},
			},
		},
		Functions:  []schema.FunctionInfo{},
		Procedures: []schema.ProcedureInfo{},
	}
	return ndcSchema, nil
}

func (c *Connector) QueryExplain(ctx context.Context, configuration *Configuration, state *State, request *schema.QueryRequest) (*schema.ExplainResponse, error) {
	return nil, schema.NotSupportedError("query explain has not been supported yet", nil)
}

func (c *Connector) MutationExplain(ctx context.Context, configuration *Configuration, state *State, request *schema.MutationRequest) (*schema.ExplainResponse, error) {
	return nil, schema.NotSupportedError("mutation explain has not been supported yet", nil)
}

func (c *Connector) Mutation(ctx context.Context, configuration *Configuration, state *State, request *schema.MutationRequest) (*schema.MutationResponse, error) {
	return nil, schema.NotSupportedError("mutation has not been supported yet", nil)
}

func (c *Connector) Query(ctx context.Context, configuration *Configuration, state *State, request *schema.QueryRequest) (schema.QueryResponse, error) {
	fmt.Printf("Configuration: %+v", configuration)
	fmt.Printf("State: %+v", state)
	fmt.Printf("Request: %+v", request)
	return nil, schema.NotSupportedError("query has not been supported yet", nil)
}
