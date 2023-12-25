import {RT_COOKIE_NAME} from '../cookies/refresh_token';
import isNonEmptyString from '../isNonEmptyString';

import {ApolloLink, FetchResult, NextLink, Observable, Operation} from '@apollo/client';


export default class AuthApolloLink extends ApolloLink {
  private bearerToken: string | null;
  constructor() {
    super();
    this.bearerToken = null;
  }

  request(operation: Operation, forward?: NextLink): Observable<FetchResult> | null {
    if (forward === undefined) {
      return null;
    }

    if (isNonEmptyString(this.bearerToken)) {
      const ctx = operation.getContext();
      const {headers} = ctx;
      operation.setContext({
        ...ctx,
        headers: {
          ...headers,
          Authorization: `Bearer ${this.bearerToken}`,
        },
      });
    }

    return forward(operation).map(data => {
      if ('extensions' in data && data.extensions != null && isNonEmptyString(data.extensions[RT_COOKIE_NAME])) {
        this.bearerToken = data.extensions[RT_COOKIE_NAME];
        delete data.extensions[RT_COOKIE_NAME];
      }
      return data;
    });
  }
}