import React, {useEffect} from "react";
import AccountsList from "./AccountsList";
import {Breadcrumbs, Typography} from "@material-ui/core";
import {useHistory} from "react-router";


export default () => {
    const history = useHistory();

    const handleRowClick = ({id}) => () => {
        history.push(`/${id}`);
    }

    useEffect(() => {
        console.log("mount");
    }, []);

    return (
        <>
            <Breadcrumbs style={{marginBottom: 32}}>
                <Typography color="textPrimary">Accounts</Typography>
            </Breadcrumbs>
            <AccountsList onRowClick={handleRowClick}/>
        </>
    )
};
