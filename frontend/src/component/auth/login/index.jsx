import * as React from 'react';
import { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { t } from 'ttag';
import { Row, Col, Card, Button } from 'antd';
import NavUtil from 'service/helper/nav_util';
import StorageUtil from 'service/helper/storage_util';
import LocaleSelect from 'component/common/locale_select.jsx';
import Form from './form';

const styles = {
    wrapper: {
        marginTop: 20
    }
};
export default function Login() {
    const navigate = useNavigate();
    const navigateTo = NavUtil.navigateTo(navigate);

    useEffect(() => {
        StorageUtil.getUserInfo() && navigateTo();
    }, []);

    function handleLogin(tenantUid) {
        const ssoUrl = `/api/v1/account/auth/sso/login/${tenantUid}`;
        window.location.href = ssoUrl;
    }

    return (
        <div>
            <div className="right content">
                <LocaleSelect />
            </div>
            <Row>
                <Col
                    xs={{ span: 24 }}
                    md={{ span: 12, offset: 6 }}
                    lg={{ span: 8, offset: 8 }}
                >
                    <Card title={t`Login`} style={styles.wrapper}>
                        <Form onChange={handleLogin} />
                    </Card>
                </Col>
            </Row>
        </div>
    );
}
