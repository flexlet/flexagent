import json
import sys
import os
import ssl
import re
import http.client
import base64
import argparse
from urllib.parse import urlparse
from urllib.parse import parse_qsl
from urllib.parse import urlencode

# error codes
ERROR_INVALID_PARAM = -1
ERROR_CONN_FAILED = 1
ERROR_SEARCH_FAILED = 2
ERROR_SEARCH_SUCCEED = 0

# api consts
PAGE_SIZE = 100
DEFAULT_HTTPS_PORT = 443
DEFAULT_HTTP_PORT = 80
URI_BASIC = "/v2"
URI_CATALOG = URI_BASIC + "/_catalog?n={limit}"
URI_TAGS_LIST = URI_BASIC + "/{repo}/tags/list"
DEFAULT_PROJECT = "library/"

def fn_parse_params():
    parser = argparse.ArgumentParser()
    parser.add_argument('--registry', help='Registry (<host>:<port>), default: REGISTRY_PYCLIENT_URL', 
        dest='registry', type=str, default=os.getenv('REGISTRY_PYCLIENT_URL'))
    parser.add_argument('--user', help='User name, default: REGISTRY_PYCLIENT_USERNAME', 
        dest='user', type=str, default=os.getenv('REGISTRY_PYCLIENT_USERNAME'))
    parser.add_argument('--password', help='Password, default: REGISTRY_PYCLIENT_PASSWORD',
         dest='password', type=str, default=os.getenv('REGISTRY_PYCLIENT_PASSWORD'))
    parser.add_argument('--insecure', help='Insecure (http), default: False',
         dest='insecure', type=bool, default=False)
    parser.add_argument('--keywords', help='Search keywords (comma seperated)', 
        dest='keywords', type=str)
    parser.add_argument('--tags', help='Search tags (comma seperated), default: all tags', 
        dest='tags', type=str, default='')
    args = parser.parse_args()
    
    if args.registry is None or args.user is None or args.password is None or args.keywords is None:
        parser.print_help()
        exit(ERROR_INVALID_PARAM)
    
    return args


def fn_search_repo(conn, headers, keywords_list):
    matched = []
    link = URI_CATALOG.format(limit = PAGE_SIZE)
    while True:
        conn.request('GET', link, headers = headers)
        resp = conn.getresponse()
        if resp.status == 200:
            body = json.loads(resp.read().decode())
            for repo in body['repositories']:
                repo_name = repo
                if repo.startswith(DEFAULT_PROJECT):
                    repo_name = repo.replace(DEFAULT_PROJECT,'')
                
                for keyword in keywords_list:
                    if keyword.lower() in repo_name.lower():
                        matched.append(repo)
                        break
        else:
            print(resp.read().decode(), file=sys.stderr)
            return None
        
        # next page
        link_header = resp.getheader('Link')
        if link_header is None:
            break
        link = link_header.replace('<','').replace('>; rel="next"','')

    return matched

def fn_list_tags(conn, headers, repo, search_tags_list):
    conn.request('GET', URI_TAGS_LIST.format(repo = repo), headers = headers)
    resp = conn.getresponse()
    if resp.status == 200:
        body = json.loads(resp.read().decode())
        if search_tags_list is None or len(search_tags_list) == 0:
            return body['tags']
        matched = []
        for tag in body['tags']:
            for search in search_tags_list:
                if search.lower() in tag.lower():
                        matched.append(tag)
                        break
        return matched
    else:
        print(resp.read().decode(), file=sys.stderr)
        return None


if __name__ == "__main__":
    args = fn_parse_params()
    if ':' in args.registry:
        item = args.registry.split(':')
        host = item[0]
        port = item[1]
    else:
        host = args.registry
        if args.insecure:
            port = DEFAULT_HTTP_PORT
        else:
            port = DEFAULT_HTTPS_PORT
    
    if args.insecure:
        conn = http.client.HTTPConnection(host, port)
    else:
        conn = http.client.HTTPSConnection(host, port, context = ssl._create_unverified_context())

    if conn is None:
        print('connect failed', file=sys.stderr)
        exit(ERROR_CONN_FAILED)

    headers = {
        'Content-Type': 'application/json;charset=UTF-8', 
        'Accept': 'application/json;charset=UTF-8'
    }

    if args.user is not None and args.password is not None:
        headers['Authorization'] = 'Basic ' + base64.b64encode('{}:{}'.format(args.user,args.password).encode()).decode()

    repos_list = fn_search_repo(conn, headers, args.keywords.split(','))

    if repos_list is None:
        print('search repo failed', file=sys.stderr)
        conn.close()
        exit(ERROR_SEARCH_FAILED)
    
    for repo in repos_list:
        tags = fn_list_tags(conn, headers, repo, args.tags.split(','))
        if tags is None:
            print('list tags failed', file=sys.stderr)
            conn.close()
            exit(ERROR_SEARCH_FAILED)
        
        repo_name = repo
        if repo.startswith(DEFAULT_PROJECT):
            repo_name = repo.replace(DEFAULT_PROJECT,'')
        
        for tag in tags:
            print('{}:{}'.format(repo_name, tag))

    conn.close()
    exit(ERROR_SEARCH_SUCCEED)