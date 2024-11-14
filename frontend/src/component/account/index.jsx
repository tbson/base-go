import * as React from "react";
import { useEffect } from 'react';
import { Outlet } from "react-router-dom";
import { useAtom } from 'jotai';
import RequestUtil from 'service/helper/request_util';
import { accountOptionSt } from './state';

const urlMap = {
    base: {
        prefix: 'account/option',
        endpoints: {
            option: ''
        }
    }
};
const urls = RequestUtil.prefixMapValues(urlMap.base);

export default function Account() {
    /*
    const [accountOption, setAccountOption] = useAtom(accountOptionSt);
    useEffect(() => {
        if (!accountOption.loaded) getOption();
    }, []);

    function getOption() {
        RequestUtil.apiCall(urls.option)
            .then((resp) => {
                setAccountOption({...resp.data, loaded: true});
            }).catch(() => {
                setAccountOption({loaded: true});
            })
    }
    */

    return <Outlet />
}

Account.displayName = "Account";
