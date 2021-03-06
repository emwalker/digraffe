import { Component, FormEvent, ChangeEvent } from 'react';
import { RelayProp } from 'react-relay';
import { Synonyms_topic as TopicType } from '__generated__/Synonyms_topic.graphql';
import { SynonymType } from 'components/types';
declare type Props = {
    relay: RelayProp;
    topic: TopicType;
};
declare type State = {
    locale: string;
    name: string;
};
declare class Synonyms extends Component<Props, State> {
    constructor(props: Props);
    onLocaleChange: (event: FormEvent<HTMLSelectElement>) => void;
    onNameChange: (event: ChangeEvent<HTMLInputElement>) => void;
    onAdd: () => void;
    onDelete: (position: number) => void;
    get synonyms(): readonly {
        readonly name: string;
        readonly locale: import("__generated__/Synonyms_topic.graphql").LocaleIdentifier;
        readonly " $fragmentRefs": import("relay-runtime").FragmentRefs<"Synonym_synonym">;
    }[];
    optimisticResponse: (synonyms: SynonymType[]) => {
        updateSynonyms: {
            alerts: never[];
            clientMutationId: null;
            topic: {
                synonyms: SynonymType[];
                displayName: string;
                id: string;
                viewerCanDeleteSynonyms: boolean;
                viewerCanUpdate: boolean;
                " $refType": "Synonyms_topic";
            };
        };
    };
    updateSynonyms: (synonyms: SynonymType[]) => void;
    renderSynonyms: () => JSX.Element;
    renderAddForm: () => JSX.Element;
    render: () => JSX.Element;
}
export declare const UnwrappedSynonyms: typeof Synonyms;
declare const _default: import("react-relay").Container<Props>;
export default _default;
