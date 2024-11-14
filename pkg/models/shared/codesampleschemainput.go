// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SchemaFile struct {
	// This field accepts []byte data or io.Reader implementations, such as *os.File.
	Content  any    `multipartForm:"content"`
	FileName string `multipartForm:"name=schema_file"`
}

func (o *SchemaFile) GetContent() any {
	if o == nil {
		return nil
	}
	return o.Content
}

func (o *SchemaFile) GetFileName() string {
	if o == nil {
		return ""
	}
	return o.FileName
}

type CodeSampleSchemaInput struct {
	// A list of languages to generate code samples for
	Languages []string `multipartForm:"name=languages"`
	// The name of the package
	PackageName *string `multipartForm:"name=package_name"`
	// The OpenAPI file to be uploaded
	SchemaFile SchemaFile `multipartForm:"file"`
	// The SDK class name
	SDKClassName *string `multipartForm:"name=sdk_class_name"`
}

func (o *CodeSampleSchemaInput) GetLanguages() []string {
	if o == nil {
		return []string{}
	}
	return o.Languages
}

func (o *CodeSampleSchemaInput) GetPackageName() *string {
	if o == nil {
		return nil
	}
	return o.PackageName
}

func (o *CodeSampleSchemaInput) GetSchemaFile() SchemaFile {
	if o == nil {
		return SchemaFile{}
	}
	return o.SchemaFile
}

func (o *CodeSampleSchemaInput) GetSDKClassName() *string {
	if o == nil {
		return nil
	}
	return o.SDKClassName
}
