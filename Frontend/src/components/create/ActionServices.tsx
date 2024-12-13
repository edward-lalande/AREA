import { Reaction, CreatePage } from "../Create";
import { AreaBox } from "../elements/AreaBox";
import { ServiceButton } from "../elements/AreaButton";
import { AreaTypography } from "../elements/AreaTypography";

type ActionServicesProps = {
    setPage: (value: CreatePage) => void;
    setSelectedActions: (value: Reaction[]) => void;
};

const ActionServices: React.FC<ActionServicesProps> = ({
    setPage,
    setSelectedActions
}) => {

    const actions = [
        {
            name: "Date Time Services",
            actions: [
                {
                    name: "At Time",
                    id: 1,
                    type: 0,
                    arguments: [
                        {
                            name: "hour",
                            type: "number"
                        },
                        {
                            name: "minute",
                            type: "number"
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

                {actions && actions.map<JSX.Element>((value) => {
                    return (<ServiceButton text={value.name} backgroundColor="#000" onClick={() => { setPage(CreatePage.ACTION); setSelectedActions(value.actions) }}/>);
                })}

            </AreaBox>

        </AreaBox>

    );

}

export default ActionServices;