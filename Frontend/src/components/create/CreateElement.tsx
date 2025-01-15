import { useEffect, useState } from "react";
import { Action, Reaction, Parameters, CreatePage } from "../Create";
import { AreaBox } from "../elements/AreaBox";
import { AddButton, AreaButton } from "../elements/AreaButton";
import { AreaTypography } from "../elements/AreaTypography";

import axios from "axios";
import { Alert, Snackbar } from "@mui/material";
import ActionServices from "./ActionServices";
import ReactionServices from "./ReactionServices";
import { useCookies } from "react-cookie";

import { User } from "../Account";

type CreateElementProps = {
    action: Action | undefined;
    actionParameters: Parameters | undefined;
    reaction: Reaction | undefined;
    reactionParameters: Parameters | undefined;
    setPage: (value: CreatePage) => void;
    setReset: (value: boolean) => void;
};

const CreateElement: React.FC<CreateElementProps> = ({
    action,
    actionParameters,
    reaction,
    reactionParameters,
    setPage,
    setReset
}) => {

    const [open, setOpen] = useState<boolean>(false);
    const [error, setError] = useState<boolean>(false);
    const [nbActions, setNbActions] = useState<number>(0);
    const [nbReactions, setNbReactions] = useState<number>(0);
  	const [cookies, setCookie] = useCookies();

    const validToken = async (action: Action, reaction: Reaction): Promise<boolean> => {

        const url: string = "http://127.0.0.1:8085/user";
        
        const token: string = cookies["token"];

        let result: boolean = true;

		await axios.get(url, {
			headers: {
				"Content-Type": "application/json",
				token
			}
		}).then((res) => {

            if (action.action_id === 2 && res.data.user.discord_token === null) {
                result = false;
            }
    
            if (action.action_id === 4 && res.data.user.github_token === null) {
                result = false;
            }
    
            if (action.action_id === 5 && res.data.user.gitlab_token === null) {
                result = false;
            }
    
            if (action.action_id === 6 && res.data.user.google_token === null) {
                result = false;
            }
    
            if (action.action_id === 9 && res.data.user.spotify_token === null) {
                result = false;
            }

            if (reaction.reaction_id === 2 && res.data.user.discord_token === null) {
                result = false;
            }

            if (reaction.reaction_id === 3 && res.data.user.dropbox_token === null) {
                result = false;
            }

            if (reaction.reaction_id === 5 && res.data.user.gitlab_token === null) {
                result = false;
            }

            if (reaction.reaction_id === 6 && res.data.user.google_token === null) {
                result = false;
            }

            if (reaction.reaction_id === 9 && res.data.user.spotify_token === null) {
                result = false;
            }

            if (reaction.reaction_id === 10 && res.data.user.asana_token === null) {
                result = false;
            }

            if (reaction.reaction_id === 14 && res.data.user.miro_token === null) {
                result = false;
            }

        });

        return result;

    }

    const createArea = async (action: Action, reaction: Reaction): Promise<boolean> => {

        let result: boolean = await validToken(action, reaction);

        if (!result) {
            setError(true);
            setReset(true);
            return false;
        }

        const url: string = "http://127.0.0.1:8080/areas";

        const data = [{
            user_token: cookies["token"],
            action: {
                name: action.name,
                action_id: action.action_id,
                action_type: action.action_type,
                ...actionParameters
            },
            reactions: [{
                name: reaction.name,
                reaction_id: reaction.reaction_id,
                reaction_type: reaction.reaction_type,
                ...reactionParameters
            }]
        }];

        axios.post(url, data).then(() => {

            setOpen(true);
            setReset(true);

        });

        return true;

    };

    const getNbActions = () => {

        const url: string = "http://127.0.0.1:8080/actions";

        axios.get(url).then((res) => {

            let actions: ActionServices[] = res.data;

            actions = actions.filter(element => element !== null);

            let count: number = 0;

            for (let i = 0; i < actions.length; i++) {
                count += actions[i].actions.length;
            }

            setNbActions(count);

        });

    }

    const getNbReactions = () => {

        const url: string = "http://127.0.0.1:8080/reactions";

        axios.get(url).then((res) => {

            let reactions: ReactionServices[] = res.data;

            reactions = reactions.filter(element => element !== null);

            let count: number = 0;

            for (let i = 0; i < reactions.length; i++) {
                count += reactions[i].reactions.length;
            }

            setNbReactions(count);

        });

    }

    useEffect(() => {

        getNbActions();
        getNbReactions();

    }, []);

    const getNbAreas = (): string => {
        return `(${nbActions} actions - ${nbReactions} reactions)`;
    }
 
    return (

        <AreaBox sx={{ height: "80vh", width: "98vw", gap: 2, mt: 10 }}>

            <AreaTypography variant="h2" text="Create an area" />
            <AreaTypography variant="h4" text={getNbAreas()} sx={{ mb: 5 }} />

            <AreaBox sx={{ flexDirection: "row", gap: 2, backgroundColor: "#000", width: "50vw", padding: "2%", borderRadius: 5 }}>
                <AreaTypography variant="h3" text={action ? "If" : "If this"} color="white" />
                {action ? <AreaTypography text={action.name} color="white" /> : <AddButton onClick={() => setPage(CreatePage.ACTION_SERVICES)}/>} 
            </AreaBox> 

            <AreaBox sx={{ flexDirection: "row", gap: 2, backgroundColor: "grey", width: "50vw", padding: "2%", borderRadius: 5, mb: 2 }}>
                <AreaTypography variant="h3" text={action ? "Then" : "Then that"} color="white" />
                {reaction ? <AreaTypography text={reaction.name} color="white" /> : action && <AddButton onClick={() => setPage(CreatePage.REACTION_SERVICES)}/>} 
            </AreaBox>

            { action && reaction && <AreaButton text="Create area" onClick={() => createArea(action, reaction)}/> }

            <Snackbar open={open} autoHideDuration={6000} onClose={() => setOpen(false)}>
                <Alert
                    onClose={() => setOpen(false)}
                    severity="success"
                    variant="filled"
                    sx={{ width:"100%", m: 2 }}
                >
                    Info: Area created!
                </Alert>
			</Snackbar>	

            <Snackbar open={error} autoHideDuration={6000} onClose={() => setError(false)}>
                <Alert
                    onClose={() => setError(false)}
                    severity="error"
                    variant="filled"
                    sx={{ width:"100%", m: 2 }}
                >
                    Error: You have to connect service to create this area!
                </Alert>
			</Snackbar>	

        </AreaBox>

    );

};

export default CreateElement;