<?php
declare(strict_types=1);

// DutchCustomerData SDK base feature

class DutchCustomerDataBaseFeature
{
    public string $version;
    public string $name;
    public bool $active;

    public function __construct()
    {
        $this->version = '0.0.1';
        $this->name = 'base';
        $this->active = true;
    }

    public function get_version(): string { return $this->version; }
    public function get_name(): string { return $this->name; }
    public function get_active(): bool { return $this->active; }

    public function init(DutchCustomerDataContext $ctx, array $options): void {}
    public function PostConstruct(DutchCustomerDataContext $ctx): void {}
    public function PostConstructEntity(DutchCustomerDataContext $ctx): void {}
    public function SetData(DutchCustomerDataContext $ctx): void {}
    public function GetData(DutchCustomerDataContext $ctx): void {}
    public function GetMatch(DutchCustomerDataContext $ctx): void {}
    public function SetMatch(DutchCustomerDataContext $ctx): void {}
    public function PrePoint(DutchCustomerDataContext $ctx): void {}
    public function PreSpec(DutchCustomerDataContext $ctx): void {}
    public function PreRequest(DutchCustomerDataContext $ctx): void {}
    public function PreResponse(DutchCustomerDataContext $ctx): void {}
    public function PreResult(DutchCustomerDataContext $ctx): void {}
    public function PreDone(DutchCustomerDataContext $ctx): void {}
    public function PreUnexpected(DutchCustomerDataContext $ctx): void {}
}
