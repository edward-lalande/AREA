import { useEffect, useState } from "react";
import { Reaction, CreatePage } from "../Create";
import { AreaBox } from "../elements/AreaBox";
import { ServiceButton } from "../elements/AreaButton";
import { AreaTypography } from "../elements/AreaTypography";
import axios from "axios";

interface ReactionServices {
    name: string;
    reactions: Reaction[];
}

type ReactionServicesProps = {
    setPage: (value: CreatePage) => void;
    setSelectedReactions: (value: Reaction[]) => void;
}

const ReactionServices: React.FC<ReactionServicesProps> = ({
    setPage,
    setSelectedReactions
}) => {

    const [reactions, setActions] = useState<ReactionServices[]>([]);

    const getReaction = () => {

        const url = "http://127.0.0.1:8080/reaction";

        axios.get(url).then((res) => {

            const actions: ReactionServices[] = res.data;

            setActions(actions.slice(0, -1));

            console.log(actions.slice(0, -1));

        });

    }

    useEffect(() => {

		getReaction();

	}, []);

    return (

        <AreaBox sx={{ height: "80vh", width: "98vw", gap: 2 }}>

            <AreaTypography variant="h2" text="Select a service" sx={{ mb: 5 }} />

            <AreaBox sx={{ height: "20vh", width: "98vw", gap: 2, flexDirection: "row" }}>

                {reactions && reactions.map<JSX.Element>((value) => {
                    return (<ServiceButton text={value.name} backgroundColor="#000" onClick={() => { setPage(CreatePage.REACTION); setSelectedReactions(value.reactions) }}/>);
                })}

            </AreaBox>

        </AreaBox>

    );

}

export default ReactionServices;