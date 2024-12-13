import { useEffect, useState } from "react";
import { Action, CreatePage } from "../Create";
import { AreaBox } from "../elements/AreaBox";
import { ServiceButton } from "../elements/AreaButton";
import { AreaTypography } from "../elements/AreaTypography";
import axios from "axios";

interface ActionServices {
    name: string;
    actions: Action[];
}

type ActionServicesProps = {
    setPage: (value: CreatePage) => void;
    setSelectedActions: (value: Action[]) => void;
};

const ActionServices: React.FC<ActionServicesProps> = ({
    setPage,
    setSelectedActions
}) => {

    const [actions, setActions] = useState<ActionServices[]>([]);

    const getActions = () => {

        const url = "http://127.0.0.1:8080/actions";

        axios.get(url).then((res) => {

            const actions: ActionServices[] = res.data;

            setActions(actions.slice(0, -1));

            console.log(actions.slice(0, -1));

        });

    }

    useEffect(() => {

		getActions();

	}, []);

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