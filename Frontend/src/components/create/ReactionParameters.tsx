import { Button } from "@mui/material";
import { Reaction, Parameters, CreatePage, Argument } from "../Create";
import { AreaBox } from "../elements/AreaBox";
import { AreaButton } from "../elements/AreaButton";
import { AreaTextField } from "../elements/AreaTextFiled";
import { AreasTypography, AreaTypography } from "../elements/AreaTypography";

type ReactionPrametersProps = {
    reaction: Reaction,
    reactionParameters: Parameters | undefined;
    setReactionParameters: (value: Parameters) => void;
    setPage: (value: CreatePage) => void;
}

const ReactionParameters: React.FC<ReactionPrametersProps> = ({
    reaction,
    reactionParameters,
    setReactionParameters,
    setPage
}) => {

    const addReactionParameter = (name: string, type: string, value: string) => {

        if (type === "number") {
            setReactionParameters(({ ...reactionParameters, [name]: Number(value) }));
        }

        if (type === "string") {
            setReactionParameters(({ ...reactionParameters, [name]: String(value) }));
        }

    };

    return (

        <AreaBox sx={{ height: "80vh", width: "98vw", gap: 2, mt: 10 }}>

            <AreaBox sx={{ width: "98vw", flexDirection: "row", justifyContent: "center", gap: 3 }}>
                <Button variant="contained" sx={{ backgroundColor: "black", color: "white", width: "5vw", borderRadius: 5, mb: 3 }} onClick={() => setPage(CreatePage.REACTION)}>Back</Button>
                <AreaTypography variant="h2" text={reaction.name} sx={{ mb: 5 }} />
            </AreaBox>

            <AreasTypography variant="h4" text={reaction.description} sx={{ mb: 5 }} />

            {  reaction.arguments && reaction.arguments.map<JSX.Element>((value: Argument) => {
                
                return (<AreaTextField label={value.display} onChange={(s) => addReactionParameter(value.name, value.type, s.target.value) } sx={{ width: "20vw" }} />);

            })}

            <AreaButton text="Validate" onClick={() => { setPage(CreatePage.CREATE); }}/>

        </AreaBox>

    );

}

export default ReactionParameters;