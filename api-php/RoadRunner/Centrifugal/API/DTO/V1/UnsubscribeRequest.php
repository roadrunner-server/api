<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: centrifugo/api/v1/api.proto

namespace RoadRunner\Centrifugal\API\DTO\V1;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>centrifugal.centrifugo.api.UnsubscribeRequest</code>
 */
class UnsubscribeRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string channel = 1 [json_name = "channel"];</code>
     */
    protected $channel = '';
    /**
     * Generated from protobuf field <code>string user = 2 [json_name = "user"];</code>
     */
    protected $user = '';
    /**
     * Generated from protobuf field <code>string client = 3 [json_name = "client"];</code>
     */
    protected $client = '';
    /**
     * Generated from protobuf field <code>string session = 4 [json_name = "session"];</code>
     */
    protected $session = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $channel
     *     @type string $user
     *     @type string $client
     *     @type string $session
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
     * Generated from protobuf field <code>string user = 2 [json_name = "user"];</code>
     * @return string
     */
    public function getUser()
    {
        return $this->user;
    }

    /**
     * Generated from protobuf field <code>string user = 2 [json_name = "user"];</code>
     * @param string $var
     * @return $this
     */
    public function setUser($var)
    {
        GPBUtil::checkString($var, True);
        $this->user = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string client = 3 [json_name = "client"];</code>
     * @return string
     */
    public function getClient()
    {
        return $this->client;
    }

    /**
     * Generated from protobuf field <code>string client = 3 [json_name = "client"];</code>
     * @param string $var
     * @return $this
     */
    public function setClient($var)
    {
        GPBUtil::checkString($var, True);
        $this->client = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string session = 4 [json_name = "session"];</code>
     * @return string
     */
    public function getSession()
    {
        return $this->session;
    }

    /**
     * Generated from protobuf field <code>string session = 4 [json_name = "session"];</code>
     * @param string $var
     * @return $this
     */
    public function setSession($var)
    {
        GPBUtil::checkString($var, True);
        $this->session = $var;

        return $this;
    }

}

