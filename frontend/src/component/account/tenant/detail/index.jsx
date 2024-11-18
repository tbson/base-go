import * as React from 'react';
import { t } from 'ttag';
import { useEffect, useState } from 'react';
import { useAtom } from 'jotai';
import { useParams } from 'react-router-dom';
import { Divider, Button } from 'antd';
import { EditOutlined } from '@ant-design/icons';
import PageHeading from 'component/common/page_heading';
import RequestUtil from 'service/helper/request_util';
import { tenantOptionSt } from '../state';
import { urls, getMessages } from '../config';
import Summary from './summary';
import Dialog from '../dialog';

export default function Tenant() {
    const { id } = useParams();
    const [item, setItem] = useState({});
    const [tenantOption, setTenantOption] = useAtom(tenantOptionSt);
    useEffect(() => {
        if (!tenantOption.loaded) {
            getOption();
        }
        getItem(id);
    }, []);

    const getOption = () => {
        return RequestUtil.apiCall(urls.option)
            .then((resp) => {
                setTenantOption({ ...resp.data, loaded: true });
            })
            .catch(() => {
                setTenantOption((prev) => ({ ...prev, loaded: true }));
            });
    };

    const getItem = (id) => {
        return RequestUtil.apiCall(`${urls.crud}${id}`).then((resp) => {
            setItem(resp.data);
        });
    };

    const onChange = (data, _id) => {
        setItem(data);
    };

    const messages = getMessages();
    return (
        <>
            <PageHeading>
                <>{messages.heading}</>
            </PageHeading>
            <Summary data={item} />
            <Divider />
            <div className="right">
                <Button
                    htmlType="button"
                    type="primary"
                    icon={<EditOutlined />}
                    onClick={() => Dialog.toggle(true, id)}
                >
                    {t`Update tenant`}
                </Button>
            </div>
            <br />
            <Dialog onChange={onChange} />
        </>
    );
}

Tenant.displayName = 'TenantDetail';