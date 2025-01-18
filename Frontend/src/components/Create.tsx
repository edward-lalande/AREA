import { useEffect, useState } from "react";
import { AreaBox } from "./elements/AreaBox";
import { AreaPaper } from "./elements/AreaPaper";

import Footer from "./Footer";
import Navbar from "./Navbar";
import CreateElement from "./create/CreateElement";
import ActionElement from "./create/ActionElement";
import ReactionElement from "./create/ReactionElement";
import ActionParameters from "./create/ActionParameters";
import ReactionParameters from "./create/ReactionParameters";
import ActionServices from "./create/ActionServices";
import ReactionServices from "./create/ReactionServices";

export enum CreatePage {
    CREATE,
    ACTION,
    REACTION,
    ACTION_SERVICES,
    REACTION_SERVICES,
    ACTION_ARGUMENTS,
    REACTION_ARGUMENTS
}

export interface Argument {
    display: string;
    name: string;
    type: string;
}

export interface Action {
    name: string;
    description: string;
    action_id: number;
    action_type: number;
    arguments: Argument[];
}

export interface Reaction {
    name: string;
    description: string;
    reaction_id: number;
    reaction_type: number;
    arguments: Argument[];
}

export type Parameters = {
    [value: string]: string | number;
}

const Create: React.FC = () => {

    const [reset, setReset] = useState<boolean>(false);

    const [action, setAction] = useState<Action>();
    const [reaction, setReaction] = useState<Reaction>();

    const [selectedActions, setSelectedActions] = useState<ActionServices | null>(null);
    const [selectedReactions, setSelectedReactions] = useState<ReactionServices | null>(null);

    const [actionParameters, setActionParameters] = useState<Parameters>();
    const [reactionParameters, setReactionParameters] = useState<Parameters>();

    const [page, setPage] = useState<CreatePage>(CreatePage.CREATE);

    useEffect(() => {

        if (reset) {
            setAction(undefined);
            setReaction(undefined);
            setSelectedActions(null);
            setSelectedReactions(null);
            setActionParameters(undefined);
            setReactionParameters(undefined);
            setReset(false);
        }

    }, [reset]);

	return (

        <AreaBox width="100vw" height="100vh" sx={{ background: "#F2F2F2" }} >

			<AreaPaper sx={{ height: "96vh", width: "98vw" }}>

                <Navbar />

                { page === CreatePage.CREATE && 
                    <CreateElement 
                        action={action}
                        reaction={reaction}
                        actionParameters={actionParameters}
                        reactionParameters={reactionParameters}
                        setPage={setPage}
                        setReset={setReset}
                    />
                }

                { page === CreatePage.ACTION_SERVICES &&
                    <ActionServices
                        setPage={setPage}
                        setSelectedActions={setSelectedActions}
                    />
                }

                { page === CreatePage.ACTION &&
                    <ActionElement
                        selectedActions={selectedActions}
                        setPage={setPage}
                        setAction={setAction}
                    />
                }

                { action && page === CreatePage.ACTION_ARGUMENTS &&
                    <ActionParameters
                        action={action}
                        actionParameters={actionParameters}
                        setActionParameters={setActionParameters}
                        setPage={setPage}
                    />
                }

                { page === CreatePage.REACTION_SERVICES &&
                    <ReactionServices
                        setPage={setPage}
                        setSelectedReactions={setSelectedReactions}
                    />
                }

                { page === CreatePage.REACTION &&

                    <ReactionElement
                        selectedReactions={selectedReactions}
                        setPage={setPage}
                        setReaction={setReaction}
                    />

                }

                { reaction && page === CreatePage.REACTION_ARGUMENTS &&
                    <ReactionParameters
                        reaction={reaction}
                        reactionParameters={reactionParameters}
                        setReactionParameters={setReactionParameters}
                        setPage={setPage}
                    />
                }

                <Footer />

			</AreaPaper>

		</AreaBox>

	);

};

export default Create;