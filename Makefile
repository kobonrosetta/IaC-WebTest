.PHONY: acctest test


test:
	@cd test && go test -v -timeout 30m

acctest:
	@echo "==> Applying acceptance testing."
	@cd terraform-provider-aws && make testacc TESTARGS='-run=TestAccDataSourceAwsEc2InstanceType_basic'
	@cd terraform-provider-aws && make testacc TESTARGS='-run=TestAccDataSourceAwsEc2InstanceType_gpu'
