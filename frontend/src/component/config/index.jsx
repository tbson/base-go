import * as React from "react";
import { useEffect } from 'react';
import { Outlet } from "react-router-dom";
import { useAtom } from 'jotai';
import RequestUtil from 'service/helper/request_util';
import { configOptionSt } from './state';

const urlMap = {
    base: {
        prefix: 'config/option',
        endpoints: {
            option: ''
        }
    }
};
const urls = RequestUtil.prefixMapValues(urlMap.base);

export default function Config() {
    const [configOption, setConfigOption] = useAtom(configOptionSt);
    useEffect(() => {
        if (!configOption.loaded) getOption();
    }, []);

    function getOption() {
        RequestUtil.apiCall(urls.option)
            .then((resp) => {
                setConfigOption({...resp.data, loaded: true});
            }).catch(() => {
                setConfigOption({loaded: true});
            })
    }

    return <Outlet />
}

Config.displayName = "Config";
