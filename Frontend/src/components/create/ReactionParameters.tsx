import { Reaction, Parameters, CreatePage, Argument } from "../Create";
import { AreaBox } from "../elements/AreaBox";
import { AreaButton } from "../elements/AreaButton";
import { AreaTextField } from "../elements/AreaTextFiled";
import { AreaTypography } from "../elements/AreaTypography";

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

            <AreaTypography variant="h2" text={reaction.name} sx={{ mb: 5 }} />

            {  reaction.arguments && reaction.arguments.map<JSX.Element>((value: Argument) => {
                
                return (<AreaTextField label={value.name} onChange={(s) => addReactionParameter(value.name, value.type, s.target.value) } sx={{ width: "20vw" }} />);

            })}

            <AreaButton text="Validate" onClick={() => { setPage(CreatePage.CREATE); }}/>

        </AreaBox>

    );

}

export default ReactionParameters;