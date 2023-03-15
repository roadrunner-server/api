<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: centrifugo/api/v1/api.proto

namespace RoadRunner\Centrifugal\API\DTO\V1;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>centrifugal.centrifugo.api.PublishRequest</code>
 */
class PublishRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string channel = 1 [json_name = "channel"];</code>
     */
    protected $channel = '';
    /**
     * Generated from protobuf field <code>bytes data = 2 [json_name = "data"];</code>
     */
    protected $data = '';
    /**
     * Generated from protobuf field <code>string b64data = 3 [json_name = "b64data"];</code>
     */
    protected $b64data = '';
    /**
     * Generated from protobuf field <code>bool skip_history = 4 [json_name = "skipHistory"];</code>
     */
    protected $skip_history = false;
    /**
     * Generated from protobuf field <code>map<string, string> tags = 5 [json_name = "tags"];</code>
     */
    private $tags;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $channel
     *     @type string $data
     *     @type string $b64data
     *     @type bool $skip_history
     *     @type array|\Google\Protobuf\Internal\MapField $tags
     * }
     */
    public function __construct($data = NULL) {
        \RoadRunner\Centrifugal\API\DTO\V1\GPBMetadata\Api::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string channel = 1 [json_name = "channel"];</code>
     * @return string
     */
    public function getChannel()
    {
        return $this->channel;
    }

    /**
     * Generated from protobuf field <code>string channel = 1 [json_name = "channel"];</code>
     * @param string $var
     * @return $this
     */
    public function setChannel($var)
    {
        GPBUtil::checkString($var, True);
        $this->channel = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>bytes data = 2 [json_name = "data"];</code>
     * @return string
     */
    public function getData()
    {
        return $this->data;
    }

    /**
     * Generated from protobuf field <code>bytes data = 2 [json_name = "data"];</code>
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
     * Generated from protobuf field <code>string b64data = 3 [json_name = "b64data"];</code>
     * @return string
     */
    public function getB64Data()
    {
        return $this->b64data;
    }

    /**
     * Generated from protobuf field <code>string b64data = 3 [json_name = "b64data"];</code>
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
     * Generated from protobuf field <code>bool skip_history = 4 [json_name = "skipHistory"];</code>
     * @return bool
     */
    public function getSkipHistory()
    {
        return $this->skip_history;
    }

    /**
     * Generated from protobuf field <code>bool skip_history = 4 [json_name = "skipHistory"];</code>
     * @param bool $var
     * @return $this
     */
    public function setSkipHistory($var)
    {
        GPBUtil::checkBool($var);
        $this->skip_history = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>map<string, string> tags = 5 [json_name = "tags"];</code>
     * @return \Google\Protobuf\Internal\MapField
     */
    public function getTags()
    {
        return $this->tags;
    }

    /**
     * Generated from protobuf field <code>map<string, string> tags = 5 [json_name = "tags"];</code>
     * @param array|\Google\Protobuf\Internal\MapField $var
     * @return $this
     */
    public function setTags($var)
    {
        $arr = GPBUtil::checkMapField($var, \Google\Protobuf\Internal\GPBType::STRING, \Google\Protobuf\Internal\GPBType::STRING);
        $this->tags = $arr;

        return $this;
    }

}

