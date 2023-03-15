<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: jobs/v1/jobs.proto

namespace RoadRunner\Jobs\DTO\V1;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Job is a main message which user might send to the RR jobs plugin
 *
 * Generated from protobuf message <code>jobs.v1.Job</code>
 */
class Job extends \Google\Protobuf\Internal\Message
{
    /**
     * job name, usually PHP class
     *
     * Generated from protobuf field <code>string job = 1 [json_name = "job"];</code>
     */
    protected $job = '';
    /**
     * unique job id
     *
     * Generated from protobuf field <code>string id = 2 [json_name = "id"];</code>
     */
    protected $id = '';
    /**
     * payload, might be embedded json or just byte-string
     *
     * Generated from protobuf field <code>string payload = 3 [json_name = "payload"];</code>
     */
    protected $payload = '';
    /**
     * headers map[string][]string
     *
     * Generated from protobuf field <code>map<string, .jobs.v1.HeaderValue> headers = 4 [json_name = "headers"];</code>
     */
    private $headers;
    /**
     * job options, contains common and driver specific fields
     *
     * Generated from protobuf field <code>.jobs.v1.Options options = 5 [json_name = "options"];</code>
     */
    protected $options = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $job
     *           job name, usually PHP class
     *     @type string $id
     *           unique job id
     *     @type string $payload
     *           payload, might be embedded json or just byte-string
     *     @type array|\Google\Protobuf\Internal\MapField $headers
     *           headers map[string][]string
     *     @type \RoadRunner\Jobs\DTO\V1\Options $options
     *           job options, contains common and driver specific fields
     * }
     */
    public function __construct($data = NULL) {
        \RoadRunner\Jobs\DTO\V1\GPBMetadata\Jobs::initOnce();
        parent::__construct($data);
    }

    /**
     * job name, usually PHP class
     *
     * Generated from protobuf field <code>string job = 1 [json_name = "job"];</code>
     * @return string
     */
    public function getJob()
    {
        return $this->job;
    }

    /**
     * job name, usually PHP class
     *
     * Generated from protobuf field <code>string job = 1 [json_name = "job"];</code>
     * @param string $var
     * @return $this
     */
    public function setJob($var)
    {
        GPBUtil::checkString($var, True);
        $this->job = $var;

        return $this;
    }

    /**
     * unique job id
     *
     * Generated from protobuf field <code>string id = 2 [json_name = "id"];</code>
     * @return string
     */
    public function getId()
    {
        return $this->id;
    }

    /**
     * unique job id
     *
     * Generated from protobuf field <code>string id = 2 [json_name = "id"];</code>
     * @param string $var
     * @return $this
     */
    public function setId($var)
    {
        GPBUtil::checkString($var, True);
        $this->id = $var;

        return $this;
    }

    /**
     * payload, might be embedded json or just byte-string
     *
     * Generated from protobuf field <code>string payload = 3 [json_name = "payload"];</code>
     * @return string
     */
    public function getPayload()
    {
        return $this->payload;
    }

    /**
     * payload, might be embedded json or just byte-string
     *
     * Generated from protobuf field <code>string payload = 3 [json_name = "payload"];</code>
     * @param string $var
     * @return $this
     */
    public function setPayload($var)
    {
        GPBUtil::checkString($var, True);
        $this->payload = $var;

        return $this;
    }

    /**
     * headers map[string][]string
     *
     * Generated from protobuf field <code>map<string, .jobs.v1.HeaderValue> headers = 4 [json_name = "headers"];</code>
     * @return \Google\Protobuf\Internal\MapField
     */
    public function getHeaders()
    {
        return $this->headers;
    }

    /**
     * headers map[string][]string
     *
     * Generated from protobuf field <code>map<string, .jobs.v1.HeaderValue> headers = 4 [json_name = "headers"];</code>
     * @param array|\Google\Protobuf\Internal\MapField $var
     * @return $this
     */
    public function setHeaders($var)
    {
        $arr = GPBUtil::checkMapField($var, \Google\Protobuf\Internal\GPBType::STRING, \Google\Protobuf\Internal\GPBType::MESSAGE, \RoadRunner\Jobs\DTO\V1\HeaderValue::class);
        $this->headers = $arr;

        return $this;
    }

    /**
     * job options, contains common and driver specific fields
     *
     * Generated from protobuf field <code>.jobs.v1.Options options = 5 [json_name = "options"];</code>
     * @return \RoadRunner\Jobs\DTO\V1\Options|null
     */
    public function getOptions()
    {
        return $this->options;
    }

    public function hasOptions()
    {
        return isset($this->options);
    }

    public function clearOptions()
    {
        unset($this->options);
    }

    /**
     * job options, contains common and driver specific fields
     *
     * Generated from protobuf field <code>.jobs.v1.Options options = 5 [json_name = "options"];</code>
     * @param \RoadRunner\Jobs\DTO\V1\Options $var
     * @return $this
     */
    public function setOptions($var)
    {
        GPBUtil::checkMessage($var, \RoadRunner\Jobs\DTO\V1\Options::class);
        $this->options = $var;

        return $this;
    }

}

