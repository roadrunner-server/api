<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: centrifugo/api/v1/api.proto

namespace RoadRunner\Centrifugal\API\DTO\V1;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>centrifugal.centrifugo.api.UserStatus</code>
 */
class UserStatus extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string user = 1 [json_name = "user"];</code>
     */
    protected $user = '';
    /**
     * Generated from protobuf field <code>int64 active = 2 [json_name = "active"];</code>
     */
    protected $active = 0;
    /**
     * Generated from protobuf field <code>int64 online = 3 [json_name = "online"];</code>
     */
    protected $online = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $user
     *     @type int|string $active
     *     @type int|string $online
     * }
     */
    public function __construct($data = NULL) {
        \RoadRunner\Centrifugal\API\DTO\V1\GPBMetadata\Api::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string user = 1 [json_name = "user"];</code>
     * @return string
     */
    public function getUser()
    {
        return $this->user;
    }

    /**
     * Generated from protobuf field <code>string user = 1 [json_name = "user"];</code>
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
     * Generated from protobuf field <code>int64 active = 2 [json_name = "active"];</code>
     * @return int|string
     */
    public function getActive()
    {
        return $this->active;
    }

    /**
     * Generated from protobuf field <code>int64 active = 2 [json_name = "active"];</code>
     * @param int|string $var
     * @return $this
     */
    public function setActive($var)
    {
        GPBUtil::checkInt64($var);
        $this->active = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int64 online = 3 [json_name = "online"];</code>
     * @return int|string
     */
    public function getOnline()
    {
        return $this->online;
    }

    /**
     * Generated from protobuf field <code>int64 online = 3 [json_name = "online"];</code>
     * @param int|string $var
     * @return $this
     */
    public function setOnline($var)
    {
        GPBUtil::checkInt64($var);
        $this->online = $var;

        return $this;
    }

}

