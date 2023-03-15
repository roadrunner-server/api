<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: websockets/v1/websockets.proto

namespace RoadRunner\Websockets\DTO\V1;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * RPC request with messages
 *
 * Generated from protobuf message <code>websockets.v1.Request</code>
 */
class Request extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>repeated .websockets.v1.Message messages = 1 [json_name = "messages"];</code>
     */
    private $messages;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type array<\RoadRunner\Websockets\DTO\V1\Message>|\Google\Protobuf\Internal\RepeatedField $messages
     * }
     */
    public function __construct($data = NULL) {
        \RoadRunner\Websockets\DTO\V1\GPBMetadata\Websockets::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>repeated .websockets.v1.Message messages = 1 [json_name = "messages"];</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getMessages()
    {
        return $this->messages;
    }

    /**
     * Generated from protobuf field <code>repeated .websockets.v1.Message messages = 1 [json_name = "messages"];</code>
     * @param array<\RoadRunner\Websockets\DTO\V1\Message>|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setMessages($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \RoadRunner\Websockets\DTO\V1\Message::class);
        $this->messages = $arr;

        return $this;
    }

}

