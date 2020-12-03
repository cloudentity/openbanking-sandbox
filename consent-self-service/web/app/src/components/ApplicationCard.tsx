import React from "react";
import {
    Accordion,
    AccordionDetails,
    AccordionSummary,
    Avatar,
    Button,
    Divider,
    Grid,
    Theme,
    Typography
} from "@material-ui/core";
import {makeStyles} from "@material-ui/core/styles";
import ExpandMoreIcon from "@material-ui/icons/ExpandMore";
import {permissionNameDescriptionMap} from "./permissionNameDescriptionMap";

const useStyles = makeStyles((theme: Theme) => ({
    expandIcon: {
        position: 'absolute', right: 32, top: 24, color: "#006580"
    }
}));

export default function ApplicationCard({client, onRevokeConsent}) {
    const classes = useStyles();

    return (
        <Accordion elevation={3} style={{marginBottom: 24}}>
            <AccordionSummary
                classes={{expandIcon: classes.expandIcon}}
                expandIcon={<Avatar><ExpandMoreIcon/></Avatar>}
                aria-controls="panel1a-content"
                id="panel1a-header"
                style={{padding: "24px 32px"}}
            >
                <div>
                    <Avatar variant={"square"} style={{
                        background: "#D5E2E5",
                        color: "#006580",
                        width: 44,
                        height: 44
                    }}>{client?.name[0]?.toUpperCase()}</Avatar>
                </div>
                <div style={{width: "100%"}}>
                    <div style={{display: "flex", alignItems: "center"}}>
                        <div style={{marginLeft: 24}}>
                            <Typography style={{color: '#C2C3C6', textTransform: 'uppercase', fontSize: 12}}>Client
                                Name</Typography>
                            <Typography
                                style={{fontSize: 16, marginTop: 4}}><strong>{client.name}</strong></Typography>
                        </div>
                    </div>
                    <div style={{marginLeft: 24, marginTop: 24}}>
                        <Typography style={{color: '#C2C3C6', textTransform: 'uppercase', fontSize: 12}}>Client
                            ID</Typography>
                        <Typography
                            style={{fontSize: 16, marginTop: 4}}><strong>{client.id}</strong></Typography>
                    </div>
                    <div style={{marginLeft: 24, marginTop: 32}}>
                        <Button
                            variant={"outlined"}
                            color={"primary"}
                            style={{textTransform: 'none'}}
                            onClick={e => {
                                e.stopPropagation();
                                window.open(client?.client_uri, "_blank")
                            }
                            }>View
                            information</Button>
                    </div>
                </div>
            </AccordionSummary>
            <AccordionDetails style={{
                flexDirection: "column",
                background: "#F4F7F8",
                padding: 0
            }}>
                {client.consents?.map(consent => (
                    <div style={{paddingBottom: 32, borderLeft: "6px solid #006580",}}
                         key={consent.consent_id}>
                        <div style={{
                            background: "#E4EEF0",
                            color: "#006580",
                            padding: '24px 0 24px 88px',
                        }}>
                            <Typography display={'block'}>
                                Consent ID: <strong>{consent.consent_id}</strong>
                            </Typography>
                            <Typography display={'block'} style={{marginTop: 16}}>
                                Granted on: <strong>{consent.creation_date_time?.split('T')[0]}</strong>
                            </Typography>
                            <Typography display={'block'} style={{marginTop: 16}}>
                                Account ID: <strong>{consent.account_ids?.join(', ')}</strong>
                            </Typography>
                        </div>
                        <div style={{paddingLeft: 88, color: "#006580"}}>
                            <Typography style={{
                                marginTop: 24,
                                textTransform: 'uppercase'
                            }}><strong>Granted permissions:</strong></Typography>
                            <Divider style={{margin: "24px 0"}}/>
                            <Grid container spacing={3}>
                                {consent.permissions?.map(permission => (
                                    <>
                                        <Grid item xs={4}
                                              style={{
                                                  textAlign: "left",
                                                  wordWrap: "break-word"
                                              }}>
                                            <Typography><strong>{permission}</strong></Typography>
                                        </Grid>
                                        <Grid item xs={8}
                                              style={{
                                                  textAlign: "left",
                                                  wordWrap: "break-word"
                                              }}>
                                            <Typography
                                                variant={"caption"}>{permissionNameDescriptionMap[permission]}</Typography>
                                        </Grid>
                                    </>
                                ))}
                            </Grid>
                            <Divider style={{margin: "24px 0"}}/>
                            <Button variant={'outlined'} color={'primary'}
                                    onClick={() => onRevokeConsent(consent)}>Revoke Access</Button>
                        </div>
                    </div>
                ))}
            </AccordionDetails>
        </Accordion>
    )
};
