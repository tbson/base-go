import * as React from "react";
import { t } from "ttag";
import { useSetRecoilState } from "recoil";
import { Form, Input } from "antd";
import FormUtil from "service/helper/form_util";
import { urls } from "../config";
import { authFlowUsernameSt, authFlowVerifIdSt } from "../states";

const formName = "ResetPwdForm";

/**
 * ResetPwdForm.
 *
 * @param {Object} object
 * @param {FormCallback} object.onChange
 */
export default function ResetPwdForm({ onChange }) {
    const setAuthFlowUsername = useSetRecoilState(authFlowUsernameSt);
    const setAuthFlowVerifId = useSetRecoilState(authFlowVerifIdSt);
    const [form] = Form.useForm();
    const initialValues = { username: "" };

    const formAttrs = {
        username: {
            name: "username",
            label: t`Email`,
            rules: [FormUtil.ruleRequired()]
        }
    };
    return (
        <Form
            name={formName}
            form={form}
            labelCol={{ span: 6 }}
            wrapperCol={{ span: 18 }}
            initialValues={{ ...initialValues }}
            onFinish={(payload) => {
                setAuthFlowUsername(payload.username);
                FormUtil.submit(urls.resetPassword, payload)
                    .then((data) => {
                        setAuthFlowVerifId(data.verif_id);
                        onChange(data);
                    })
                    .catch(FormUtil.setFormErrors(form));
            }}
        >
            <Form.Item {...formAttrs.username}>
                <Input autoFocus />
            </Form.Item>
        </Form>
    );
}

ResetPwdForm.displayName = formName;
ResetPwdForm.formName = formName;
