<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: websockets/v1/websockets.proto

namespace RoadRunner\Websockets\DTO\V1;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>websockets.v1.Message</code>
 */
class Message extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string command = 1 [json_name = "command"];</code>
     */
    protected $command = '';
    /**
     * Generated from protobuf field <code>repeated string topics = 2 [json_name = "topics"];</code>
     */
    private $topics;
    /**
     * Generated from protobuf field <code>bytes payload = 3 [json_name = "payload"];</code>
     */
    protected $payload = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $command
     *     @type array<string>|\Google\Protobuf\Internal\RepeatedField $topics
     *     @type string $payload
     * }
     */
    public function __construct($data = NULL) {
        \RoadRunner\Websockets\DTO\V1\GPBMetadata\Websockets::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string command = 1 [json_name = "command"];</code>
     * @return string
     */
    public function getCommand()
    {
        return $this->command;
    }

    /**
     * Generated from protobuf field <code>string command = 1 [json_name = "command"];</code>
     * @param string $var
     * @return $this
     */
    public function setCommand($var)
    {
        GPBUtil::checkString($var, True);
        $this->command = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated string topics = 2 [json_name = "topics"];</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getTopics()
    {
        return $this->topics;
    }

    /**
     * Generated from protobuf field <code>repeated string topics = 2 [json_name = "topics"];</code>
     * @param array<string>|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setTopics($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->topics = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>bytes payload = 3 [json_name = "payload"];</code>
     * @return string
     */
    public function getPayload()
    {
        return $this->payload;
    }

    /**
     * Generated from protobuf field <code>bytes payload = 3 [json_name = "payload"];</code>
     * @param string $var
     * @return $this
     */
    public function setPayload($var)
    {
        GPBUtil::checkString($var, False);
        $this->payload = $var;

        return $this;
    }

}

