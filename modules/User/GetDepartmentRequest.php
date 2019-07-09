<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: user/department.proto

namespace User;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * user info 请求数据
 *
 * Generated from protobuf message <code>user.GetDepartmentRequest</code>
 */
final class GetDepartmentRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string DepId = 1;</code>
     */
    private $DepId = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $DepId
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\User\Department::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string DepId = 1;</code>
     * @return string
     */
    public function getDepId()
    {
        return $this->DepId;
    }

    /**
     * Generated from protobuf field <code>string DepId = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setDepId($var)
    {
        GPBUtil::checkString($var, True);
        $this->DepId = $var;

        return $this;
    }

}
