import React from "react";
import ClientsList from "./ClientsList";
import {Breadcrumbs, Typography} from "@material-ui/core";


export default ({clients, onRevokeClient, onRevokeConsent}) => {
    return (
        <>
            <Breadcrumbs style={{marginBottom: 32}}>
                <Typography color="textPrimary">Clients</Typography>
            </Breadcrumbs>
            <ClientsList clients={clients} onRevokeClient={onRevokeClient} onRevokeConsent={onRevokeConsent}/>
        </>
    )
};
