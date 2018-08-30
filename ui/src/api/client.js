import axios from 'axios';

function merge(obj1, obj2) {
    var obj3 = {};
    for (var attrname in obj1) {
        obj3[attrname] = obj1[attrname];
    }
    for (var attrname in obj2) {
        obj3[attrname] = obj2[attrname];
    }
    return obj3;
}
function client (scheme, host, timeout) {
    scheme = scheme || 'http';
    host = host || 'localhost:8080';
    timeout = timeout || 20000;
    var client = axios;
    var urlPrefix = scheme + '://' + host;
    client.createMemos = function (path, data, config) {
        var cfg = {
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
    client.deleteMemos = function (path, config) {
        var cfg = {
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
    client.listMemos = function (path, config) {
        var cfg = {
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
    client.showMemos = function (path, config) {
        var cfg = {
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
    client.updateMemos = function (path, data, config) {
        var cfg = {
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
