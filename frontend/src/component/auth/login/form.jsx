import * as React from 'react';
import { Button, Row, Col, Form, Input } from 'antd';
import { t } from 'ttag';
import { CheckOutlined } from '@ant-design/icons';
import FormUtil from 'service/helper/form_util';

const formName = 'LoginForm';

export default function LoginForm({ onChange, children }) {
    const [form] = Form.useForm();
    const initialValues = {
        tenantUid: ''
    };

    const formAttrs = {
        tenantUid: {
            name: 'tenantUid',
            label: t`Company code`,
            rules: [FormUtil.ruleRequired()]
        }
    };

    return (
        <Form
            form={form}
            labelCol={{ span: 8 }}
            wrapperCol={{ span: 16 }}
            initialValues={{ ...initialValues }}
            onFinish={(payload) => {
                onChange(payload.tenantUid);
            }}
        >
            <Form.Item {...formAttrs.tenantUid}>
                <Input autoFocus />
            </Form.Item>

            <br />
            <Row>
                <Col span={12}>{children}</Col>
                <Col span={12} className="right">
                    <Button type="primary" htmlType="submit" icon={<CheckOutlined />}>
                        {t`Login`}
                    </Button>
                </Col>
            </Row>
        </Form>
    );
}
LoginForm.displayName = formName;
LoginForm.formName = formName;
