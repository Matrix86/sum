<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/sum.proto

namespace Sum;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>sum.Call</code>
 */
class Call extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>uint64 oracle_id = 1;</code>
     */
    private $oracle_id = 0;
    /**
     * Generated from protobuf field <code>repeated string args = 2;</code>
     */
    private $args;

    public function __construct() {
        \GPBMetadata\Proto\Sum::initOnce();
        parent::__construct();
    }

    /**
     * Generated from protobuf field <code>uint64 oracle_id = 1;</code>
     * @return int|string
     */
    public function getOracleId()
    {
        return $this->oracle_id;
    }

    /**
     * Generated from protobuf field <code>uint64 oracle_id = 1;</code>
     * @param int|string $var
     * @return $this
     */
    public function setOracleId($var)
    {
        GPBUtil::checkUint64($var);
        $this->oracle_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated string args = 2;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getArgs()
    {
        return $this->args;
    }

    /**
     * Generated from protobuf field <code>repeated string args = 2;</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setArgs($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->args = $arr;

        return $this;
    }

}

