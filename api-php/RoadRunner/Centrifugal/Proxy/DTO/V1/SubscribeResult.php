<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: centrifugo/proxy/v1/proxy.proto

namespace RoadRunner\Centrifugal\Proxy\DTO\V1;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>centrifugal.centrifugo.proxy.SubscribeResult</code>
 */
class SubscribeResult extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>int64 expire_at = 1 [json_name = "expireAt"];</code>
     */
    protected $expire_at = 0;
    /**
     * Generated from protobuf field <code>bytes info = 2 [json_name = "info"];</code>
     */
    protected $info = '';
    /**
     * Generated from protobuf field <code>string b64info = 3 [json_name = "b64info"];</code>
     */
    protected $b64info = '';
    /**
     * Generated from protobuf field <code>bytes data = 4 [json_name = "data"];</code>
     */
    protected $data = '';
    /**
     * Generated from protobuf field <code>string b64data = 5 [json_name = "b64data"];</code>
     */
    protected $b64data = '';
    /**
     * Generated from protobuf field <code>.centrifugal.centrifugo.proxy.SubscribeOptionOverride override = 6 [json_name = "override"];</code>
     */
    protected $override = null;
    /**
     * Generated from protobuf field <code>repeated string allow = 7 [json_name = "allow"];</code>
     */
    private $allow;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int|string $expire_at
     *     @type string $info
     *     @type string $b64info
     *     @type string $data
     *     @type string $b64data
     *     @type \RoadRunner\Centrifugal\Proxy\DTO\V1\SubscribeOptionOverride $override
     *     @type array<string>|\Google\Protobuf\Internal\RepeatedField $allow
     * }
     */
    public function __construct($data = NULL) {
        \RoadRunner\Centrifugal\Proxy\DTO\V1\GPBMetadata\Proxy::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>int64 expire_at = 1 [json_name = "expireAt"];</code>
     * @return int|string
     */
    public function getExpireAt()
    {
        return $this->expire_at;
    }

    /**
     * Generated from protobuf field <code>int64 expire_at = 1 [json_name = "expireAt"];</code>
     * @param int|string $var
     * @return $this
     */
    public function setExpireAt($var)
    {
        GPBUtil::checkInt64($var);
        $this->expire_at = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>bytes info = 2 [json_name = "info"];</code>
     * @return string
     */
    public function getInfo()
    {
        return $this->info;
    }

    /**
     * Generated from protobuf field <code>bytes info = 2 [json_name = "info"];</code>
     * @param string $var
     * @return $this
     */
    public function setInfo($var)
    {
        GPBUtil::checkString($var, False);
        $this->info = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string b64info = 3 [json_name = "b64info"];</code>
     * @return string
     */
    public function getB64Info()
    {
        return $this->b64info;
    }

    /**
     * Generated from protobuf field <code>string b64info = 3 [json_name = "b64info"];</code>
     * @param string $var
     * @return $this
     */
    public function setB64Info($var)
    {
        GPBUtil::checkString($var, True);
        $this->b64info = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>bytes data = 4 [json_name = "data"];</code>
     * @return string
     */
    public function getData()
    {
        return $this->data;
    }

    /**
     * Generated from protobuf field <code>bytes data = 4 [json_name = "data"];</code>
     * @param string $var
     * @return $this
     */
    public function setData($var)
    {
        GPBUtil::checkString($var, False);
        $this->data = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string b64data = 5 [json_name = "b64data"];</code>
     * @return string
     */
    public function getB64Data()
    {
        return $this->b64data;
    }

    /**
     * Generated from protobuf field <code>string b64data = 5 [json_name = "b64data"];</code>
     * @param string $var
     * @return $this
     */
    public function setB64Data($var)
    {
        GPBUtil::checkString($var, True);
        $this->b64data = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.centrifugal.centrifugo.proxy.SubscribeOptionOverride override = 6 [json_name = "override"];</code>
     * @return \RoadRunner\Centrifugal\Proxy\DTO\V1\SubscribeOptionOverride|null
     */
    public function getOverride()
    {
        return $this->override;
    }

    public function hasOverride()
    {
        return isset($this->override);
    }

    public function clearOverride()
    {
        unset($this->override);
    }

    /**
     * Generated from protobuf field <code>.centrifugal.centrifugo.proxy.SubscribeOptionOverride override = 6 [json_name = "override"];</code>
     * @param \RoadRunner\Centrifugal\Proxy\DTO\V1\SubscribeOptionOverride $var
     * @return $this
     */
    public function setOverride($var)
    {
        GPBUtil::checkMessage($var, \RoadRunner\Centrifugal\Proxy\DTO\V1\SubscribeOptionOverride::class);
        $this->override = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated string allow = 7 [json_name = "allow"];</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getAllow()
    {
        return $this->allow;
    }

    /**
     * Generated from protobuf field <code>repeated string allow = 7 [json_name = "allow"];</code>
     * @param array<string>|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setAllow($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->allow = $arr;

        return $this;
    }

}

