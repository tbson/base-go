import * as React from "react";
import { t } from "ttag";
import { Form, Input } from "antd";
import FormUtil from "service/helper/form_util";
import { urls } from "../../config";

const formName = "UpdateProfileForm";

export default function UpdateProfileForm({ data, onChange }) {
    const [form] = Form.useForm();

    const formAttrs = {
        mobile: {
            name: "mobile",
            label: t`Mobile`,
            rules: [FormUtil.ruleRequired()]
        },
        first_name: {
            name: "first_name",
            label: t`First Name`,
            rules: [FormUtil.ruleRequired()]
        },
        last_name: {
            name: "last_name",
            label: t`Last Name`,
            rules: [FormUtil.ruleRequired()]
        }
    };

    return (
        <Form
            form={form}
            name={formName}
            labelCol={{ span: 6 }}
            wrapperCol={{ span: 18 }}
            initialValues={{ ...data }}
            onFinish={(payload) =>
                FormUtil.submit(urls.profile, payload, "put")
                    .then((data) => onChange(data))
                    .catch(FormUtil.setFormErrors(form))
            }
        >
            <Form.Item {...formAttrs.mobile}>
                <Input autoFocus />
            </Form.Item>
            <Form.Item {...formAttrs.first_name}>
                <Input />
            </Form.Item>
            <Form.Item {...formAttrs.last_name}>
                <Input />
            </Form.Item>
        </Form>
    );
}

UpdateProfileForm.displayName = formName;
UpdateProfileForm.formName = formName;
