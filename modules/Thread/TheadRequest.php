<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: thread/thread.proto

namespace Thread;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * thread 请求数据
 *
 * Generated from protobuf message <code>thread.TheadRequest</code>
 */
final class TheadRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string Id = 1;</code>
     */
    private $Id = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $Id
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Thread\Thread::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string Id = 1;</code>
     * @return string
     */
    public function getId()
    {
        return $this->Id;
    }

    /**
     * Generated from protobuf field <code>string Id = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setId($var)
    {
        GPBUtil::checkString($var, True);
        $this->Id = $var;

        return $this;
    }

}

