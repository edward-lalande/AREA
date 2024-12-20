import { Reaction, CreatePage } from "../Create";
import { AreaBox } from "../elements/AreaBox";
import { ServiceButton } from "../elements/AreaButton";
import { AreaTypography } from "../elements/AreaTypography";


type ReactionElementProps = {
    selectedReactions: Reaction[];
    setPage: (value: CreatePage) => void;
    setReaction: (value: Reaction) => void;
}

const ReactionElement: React.FC<ReactionElementProps> = ({
    selectedReactions,
    setPage,
    setReaction
}) => {

    return (

        <AreaBox sx={{ height: "80vh", width: "98vw", gap: 2 }}>

            <AreaTypography variant="h2" text="Select a reaction" sx={{ mb: 5 }} />

            <AreaBox sx={{ height: "20vh", width: "98vw", gap: 2, flexDirection: "row" }}>

                {selectedReactions && selectedReactions.map<JSX.Element>((value: Reaction) => {
                    return (<ServiceButton text={value.name} backgroundColor="#000" onClick={() => {setPage(CreatePage.REACTION_ARGUMENTS); setReaction(value)}}/>);
                })}

            </AreaBox>

        </AreaBox>

    );

}

export default ReactionElement;