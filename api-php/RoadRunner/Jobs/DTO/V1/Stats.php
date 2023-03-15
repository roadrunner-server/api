<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: jobs/v1/jobs.proto

namespace RoadRunner\Jobs\DTO\V1;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>jobs.v1.Stats</code>
 */
class Stats extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>repeated .jobs.v1.Stat stats = 1 [json_name = "stats"];</code>
     */
    private $stats;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type array<\RoadRunner\Jobs\DTO\V1\Stat>|\Google\Protobuf\Internal\RepeatedField $stats
     * }
     */
    public function __construct($data = NULL) {
        \RoadRunner\Jobs\DTO\V1\GPBMetadata\Jobs::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>repeated .jobs.v1.Stat stats = 1 [json_name = "stats"];</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getStats()
    {
        return $this->stats;
    }

    /**
     * Generated from protobuf field <code>repeated .jobs.v1.Stat stats = 1 [json_name = "stats"];</code>
     * @param array<\RoadRunner\Jobs\DTO\V1\Stat>|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setStats($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \RoadRunner\Jobs\DTO\V1\Stat::class);
        $this->stats = $arr;

        return $this;
    }

}

