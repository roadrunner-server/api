<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: http/v1beta/http.proto

namespace RoadRunner\HTTP\DTO\V1BETA1;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>http.v1beta.Response</code>
 */
class Response extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>int64 status = 1 [json_name = "status"];</code>
     */
    protected $status = 0;
    /**
     * Generated from protobuf field <code>map<string, .http.v1beta.HeaderValue> headers = 2 [json_name = "headers"];</code>
     */
    private $headers;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int|string $status
     *     @type array|\Google\Protobuf\Internal\MapField $headers
     * }
     */
    public function __construct($data = NULL) {
        \RoadRunner\HTTP\DTO\V1BETA1\GPBMetadata\Http::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>int64 status = 1 [json_name = "status"];</code>
     * @return int|string
     */
    public function getStatus()
    {
        return $this->status;
    }

    /**
     * Generated from protobuf field <code>int64 status = 1 [json_name = "status"];</code>
     * @param int|string $var
     * @return $this
     */
    public function setStatus($var)
    {
        GPBUtil::checkInt64($var);
        $this->status = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>map<string, .http.v1beta.HeaderValue> headers = 2 [json_name = "headers"];</code>
     * @return \Google\Protobuf\Internal\MapField
     */
    public function getHeaders()
    {
        return $this->headers;
    }

    /**
     * Generated from protobuf field <code>map<string, .http.v1beta.HeaderValue> headers = 2 [json_name = "headers"];</code>
     * @param array|\Google\Protobuf\Internal\MapField $var
     * @return $this
     */
    public function setHeaders($var)
    {
        $arr = GPBUtil::checkMapField($var, \Google\Protobuf\Internal\GPBType::STRING, \Google\Protobuf\Internal\GPBType::MESSAGE, \RoadRunner\HTTP\DTO\V1BETA1\HeaderValue::class);
        $this->headers = $arr;

        return $this;
    }

}

