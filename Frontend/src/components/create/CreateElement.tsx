import { Action, Reaction, Parameters, CreatePage } from "../Create";
import { AreaBox } from "../elements/AreaBox";
import { AddButton, AreaButton } from "../elements/AreaButton";
import { AreaTypography } from "../elements/AreaTypography";

import axios from "axios";

type CreateElementProps = {
    action: Action | undefined;
    actionParameters: Parameters | undefined;
    reaction: Reaction | undefined;
    reactionParameters: Parameters | undefined;
    setPage: (value: CreatePage) => void;
};

const CreateElement: React.FC<CreateElementProps> = ({
    action,
    actionParameters,
    reaction,
    reactionParameters,
    setPage,
}) => {

    const createArea = (action: Action, reaction: Reaction) => {

        const url: string = "http://127.0.0.1:8080/area";

        const data = [{
            user_token: "AREA",
            action: {
                action_id: action.id,
                action_type: action.type,
                ...actionParameters
            },
            reactions: [{
                reaction_id: reaction.id,
                reaction_type: reaction.type,
                ...reactionParameters
            }]
        }];

        console.log(data);

        axios.post(url, data).then(() => {

            //window.location.href = "/";

        });

    };

    return (

        <AreaBox sx={{ height: "80vh", width: "98vw", gap: 2, mt: 10 }}>

            <AreaTypography variant="h2" text="Create an area" sx={{ mb: 5 }} />

            <AreaBox sx={{ flexDirection: "row", gap: 2, backgroundColor: "#000", width: "50vw", padding: "2%", borderRadius: 5 }}>
                <AreaTypography variant="h3" text={action ? "If" : "If this"} color="white" />
                {action ? <AreaTypography text={action.name} color="white" /> : <AddButton onClick={() => setPage(CreatePage.ACTION_SERVICES)}/>} 
            </AreaBox> 

            <AreaBox sx={{ flexDirection: "row", gap: 2, backgroundColor: "grey", width: "50vw", padding: "2%", borderRadius: 5, mb: 2 }}>
                <AreaTypography variant="h3" text={action ? "Then" : "Then that"} color="white" />
                {reaction ? <AreaTypography text={reaction.name} color="white" /> : action && <AddButton onClick={() => setPage(CreatePage.REACTION_SERVICES)}/>} 
            </AreaBox>

            { action && reaction && <AreaButton text="Create area" onClick={() => createArea(action, reaction)}/> }

        </AreaBox>

    );

};

export default CreateElement;