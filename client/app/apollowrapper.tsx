'use client'
import { ApolloLink, HttpLink } from "@apollo/client";
import {ApolloNextAppProvider, NextSSRInMemoryCache, SSRMultipartLink, NextSSRApolloClient} from "@apollo/experimental-nextjs-app-support/ssr";

function makeClient() {
  const httpLink = new HttpLink({
    uri: "http://localhost:3001/",
    headers: {
      accept: 'application/json',
    }
  });

  return new NextSSRApolloClient({
    cache : new NextSSRInMemoryCache(),
    link:
      typeof window === "undefined"
        ? ApolloLink.from([
            new SSRMultipartLink({
              stripDefer: true,
            }),
            httpLink,
          ])
        : httpLink,
  });
}

// you need to create a component to wrap your app in
export function ApolloWrapper({ children }: React.PropsWithChildren) {
  return (
    <ApolloNextAppProvider makeClient={makeClient}>
      {children}
    </ApolloNextAppProvider>
  );
}