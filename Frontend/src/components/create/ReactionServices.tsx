import { Reaction, CreatePage } from "../Create";
import { AreaBox } from "../elements/AreaBox";
import { ServiceButton } from "../elements/AreaButton";
import { AreaTypography } from "../elements/AreaTypography";

type ReactionServicesProps = {
    setPage: (value: CreatePage) => void;
    setSelectedReactions: (value: Reaction[]) => void;
}

const ReactionServices: React.FC<ReactionServicesProps> = ({
    setPage,
    setSelectedReactions
}) => {

    const reactions = [
        {
            name: "Discord Services",
            reactions: [
                {
                    name: "Send message on channel",
                    id: 2,
                    type: 0,
                    arguments: [
                        {
                            name: "channel_id",
                            type: "string"
                        },
                        {
                            name: "message",
                            type: "string"
                        }
                    ]
                }
            ]
        }
    ];

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