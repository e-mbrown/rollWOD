import { ApolloClient, HttpLink, InMemoryCache, gql } from "@apollo/client";
import { registerApolloClient } from "@apollo/experimental-nextjs-app-support/rsc";

export const { getClient } = registerApolloClient(() => {
  return new ApolloClient({
    cache: new InMemoryCache(),
    link: new HttpLink({
      uri: "http://localhost:8080",
      headers: {
        accept: 'application/json',
      } 
    }),
  });
});