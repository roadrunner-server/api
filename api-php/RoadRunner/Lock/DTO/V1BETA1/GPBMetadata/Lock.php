<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: lock/v1beta1/lock.proto

namespace RoadRunner\Lock\DTO\V1BETA1\GPBMetadata;

class Lock
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        $pool->internalAddGeneratedFile(
            '
�
lock/v1beta1/lock.protolock.v1beta1"v
Request
resource (	Rresource
id (	Rid
ttl (H Rttl�
wait (HRwait�B
_ttlB
_wait"
Response
ok (RokBVZlock/v1beta1�RoadRunner\\Lock\\DTO\\V1BETA1�\'RoadRunner\\Lock\\DTO\\V1BETA1\\GPBMetadatabproto3'
        , true);

        static::$is_initialized = true;
    }
}

