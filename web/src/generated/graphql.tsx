import { gql } from '@apollo/client';
import * as Apollo from '@apollo/client';
export type Maybe<T> = T | null;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
};

export type Game = {
  __typename?: 'Game';
  id: Scalars['ID'];
  letters: Scalars['String'];
  correct: Array<Scalars['String']>;
  incorrect: Array<Scalars['String']>;
  total: Scalars['Int'];
};

export type GuessResult = {
  __typename?: 'GuessResult';
  correct: Scalars['Boolean'];
  word: Scalars['String'];
  game: Game;
};

export type Query = {
  __typename?: 'Query';
  game: Game;
};


export type QueryGameArgs = {
  id: Scalars['ID'];
};

export type Mutation = {
  __typename?: 'Mutation';
  newGame: Game;
  addGuess: GuessResult;
};


export type MutationAddGuessArgs = {
  id: Scalars['ID'];
  guess: Scalars['String'];
};

export type Subscription = {
  __typename?: 'Subscription';
  watchGame: GuessResult;
};


export type SubscriptionWatchGameArgs = {
  id: Scalars['ID'];
};

export type UpdateFieldsFragment = (
  { __typename?: 'Game' }
  & Pick<Game, 'id' | 'letters' | 'correct' | 'incorrect'>
);

export type NewGameMutationVariables = Exact<{ [key: string]: never; }>;


export type NewGameMutation = (
  { __typename?: 'Mutation' }
  & { newGame: (
    { __typename?: 'Game' }
    & Pick<Game, 'total'>
    & UpdateFieldsFragment
  ) }
);

export type CurrGameQueryVariables = Exact<{
  id: Scalars['ID'];
}>;


export type CurrGameQuery = (
  { __typename?: 'Query' }
  & { game: (
    { __typename?: 'Game' }
    & Pick<Game, 'total'>
    & UpdateFieldsFragment
  ) }
);

export type AddGuessMutationVariables = Exact<{
  id: Scalars['ID'];
  guess: Scalars['String'];
}>;


export type AddGuessMutation = (
  { __typename?: 'Mutation' }
  & { addGuess: (
    { __typename?: 'GuessResult' }
    & Pick<GuessResult, 'correct' | 'word'>
    & { game: (
      { __typename?: 'Game' }
      & UpdateFieldsFragment
    ) }
  ) }
);

export type WatchGameSubscriptionVariables = Exact<{
  id: Scalars['ID'];
}>;


export type WatchGameSubscription = (
  { __typename?: 'Subscription' }
  & { watchGame: (
    { __typename?: 'GuessResult' }
    & Pick<GuessResult, 'correct' | 'word'>
    & { game: (
      { __typename?: 'Game' }
      & UpdateFieldsFragment
    ) }
  ) }
);

export const UpdateFieldsFragmentDoc = gql`
    fragment updateFields on Game {
  id
  letters
  correct
  incorrect
}
    `;
export const NewGameDocument = gql`
    mutation newGame {
  newGame {
    ...updateFields
    total
  }
}
    ${UpdateFieldsFragmentDoc}`;
export type NewGameMutationFn = Apollo.MutationFunction<NewGameMutation, NewGameMutationVariables>;

/**
 * __useNewGameMutation__
 *
 * To run a mutation, you first call `useNewGameMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useNewGameMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [newGameMutation, { data, loading, error }] = useNewGameMutation({
 *   variables: {
 *   },
 * });
 */
export function useNewGameMutation(baseOptions?: Apollo.MutationHookOptions<NewGameMutation, NewGameMutationVariables>) {
        return Apollo.useMutation<NewGameMutation, NewGameMutationVariables>(NewGameDocument, baseOptions);
      }
export type NewGameMutationHookResult = ReturnType<typeof useNewGameMutation>;
export type NewGameMutationResult = Apollo.MutationResult<NewGameMutation>;
export type NewGameMutationOptions = Apollo.BaseMutationOptions<NewGameMutation, NewGameMutationVariables>;
export const CurrGameDocument = gql`
    query currGame($id: ID!) {
  game(id: $id) {
    ...updateFields
    total
  }
}
    ${UpdateFieldsFragmentDoc}`;

/**
 * __useCurrGameQuery__
 *
 * To run a query within a React component, call `useCurrGameQuery` and pass it any options that fit your needs.
 * When your component renders, `useCurrGameQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useCurrGameQuery({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useCurrGameQuery(baseOptions: Apollo.QueryHookOptions<CurrGameQuery, CurrGameQueryVariables>) {
        return Apollo.useQuery<CurrGameQuery, CurrGameQueryVariables>(CurrGameDocument, baseOptions);
      }
export function useCurrGameLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<CurrGameQuery, CurrGameQueryVariables>) {
          return Apollo.useLazyQuery<CurrGameQuery, CurrGameQueryVariables>(CurrGameDocument, baseOptions);
        }
export type CurrGameQueryHookResult = ReturnType<typeof useCurrGameQuery>;
export type CurrGameLazyQueryHookResult = ReturnType<typeof useCurrGameLazyQuery>;
export type CurrGameQueryResult = Apollo.QueryResult<CurrGameQuery, CurrGameQueryVariables>;
export const AddGuessDocument = gql`
    mutation addGuess($id: ID!, $guess: String!) {
  addGuess(id: $id, guess: $guess) {
    correct
    word
    game {
      ...updateFields
    }
  }
}
    ${UpdateFieldsFragmentDoc}`;
export type AddGuessMutationFn = Apollo.MutationFunction<AddGuessMutation, AddGuessMutationVariables>;

/**
 * __useAddGuessMutation__
 *
 * To run a mutation, you first call `useAddGuessMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useAddGuessMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [addGuessMutation, { data, loading, error }] = useAddGuessMutation({
 *   variables: {
 *      id: // value for 'id'
 *      guess: // value for 'guess'
 *   },
 * });
 */
export function useAddGuessMutation(baseOptions?: Apollo.MutationHookOptions<AddGuessMutation, AddGuessMutationVariables>) {
        return Apollo.useMutation<AddGuessMutation, AddGuessMutationVariables>(AddGuessDocument, baseOptions);
      }
export type AddGuessMutationHookResult = ReturnType<typeof useAddGuessMutation>;
export type AddGuessMutationResult = Apollo.MutationResult<AddGuessMutation>;
export type AddGuessMutationOptions = Apollo.BaseMutationOptions<AddGuessMutation, AddGuessMutationVariables>;
export const WatchGameDocument = gql`
    subscription watchGame($id: ID!) {
  watchGame(id: $id) {
    correct
    word
    game {
      ...updateFields
    }
  }
}
    ${UpdateFieldsFragmentDoc}`;

/**
 * __useWatchGameSubscription__
 *
 * To run a query within a React component, call `useWatchGameSubscription` and pass it any options that fit your needs.
 * When your component renders, `useWatchGameSubscription` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the subscription, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useWatchGameSubscription({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useWatchGameSubscription(baseOptions: Apollo.SubscriptionHookOptions<WatchGameSubscription, WatchGameSubscriptionVariables>) {
        return Apollo.useSubscription<WatchGameSubscription, WatchGameSubscriptionVariables>(WatchGameDocument, baseOptions);
      }
export type WatchGameSubscriptionHookResult = ReturnType<typeof useWatchGameSubscription>;
export type WatchGameSubscriptionResult = Apollo.SubscriptionResult<WatchGameSubscription>;