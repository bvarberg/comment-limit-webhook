aws cloudformation package \
  --template-file $TEMPLATE_FILE \
  --s3-bucket $S3_BUCKET \
  --output-template-file $OUTPUT_TEMPLATE_FILE
