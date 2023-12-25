import AuthApolloLink from './AuthApolloLink';
import createCacheAndRestoreOnce from './createCacheAndRestoreOnce';

import {ApolloClient, NormalizedCacheObject, concat, createHttpLink} from '@apollo/client';
import fetch from 'cross-fetch';

export default function createAuthApolloClient(
  graphqlUri: string,
): ApolloClient<NormalizedCacheObject> {
  const authLink = new AuthApolloLink();
  const httpLink = createHttpLink({
    uri: graphqlUri,
    fetch,
  });
  return new ApolloClient({
    cache: createCacheAndRestoreOnce(),
    link: concat(httpLink, authLink),
  });
}