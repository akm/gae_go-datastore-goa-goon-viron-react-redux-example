import axios, { AxiosRequestConfig } from 'axios';
import * as objectAssign from 'object-assign';
import { Memo } from '../states/Memo'

const merge = (...args) => objectAssign({}, ...args)

function client (scheme = 'http', host = 'localhost:8080', timeout = 20000) {
    let client = axios;
    let urlPrefix = scheme + '://' + host;
    client.createMemos = function (path: string, data: AxiosRequestConfig, config?: AxiosRequestConfig) {
        let cfg: AxiosRequestConfig = {
            timeout: timeout,
            url: urlPrefix + path,
            method: 'post',
            data: data,
            responseType: 'json'
        };
        if (config) {
            cfg = merge(cfg, config);
        }
        return client(cfg);
    };
    client.deleteMemos = function (path: string, config?: AxiosRequestConfig) {
        let cfg: AxiosRequestConfig = {
            timeout: timeout,
            url: urlPrefix + path,
            method: 'delete',
            responseType: 'json'
        };
        if (config) {
            cfg = merge(cfg, config);
        }
        return client(cfg);
    };
    client.listMemos = function (path: string, config?: AxiosRequestConfig) {
        let cfg: AxiosRequestConfig = {
            timeout: timeout,
            url: urlPrefix + path,
            method: 'get',
            responseType: 'json'
        };
        if (config) {
            cfg = merge(cfg, config);
        }
        return client(cfg);
    };
    client.showMemos = function (path: string, config?: AxiosRequestConfig) {
        let cfg: AxiosRequestConfig = {
            timeout: timeout,
            url: urlPrefix + path,
            method: 'get',
            responseType: 'json'
        };
        if (config) {
            cfg = merge(cfg, config);
        }
        return client(cfg);
    };
    client.updateMemos = function (path: string, data: Memo, config?: AxiosRequestConfig) {
        let cfg: AxiosRequestConfig = {
            timeout: timeout,
            url: urlPrefix + path,
            method: 'put',
            data: data,
            responseType: 'json'
        };
        if (config) {
            cfg = merge(cfg, config);
        }
        return client(cfg);
    };
    return client;
}

export default client;
