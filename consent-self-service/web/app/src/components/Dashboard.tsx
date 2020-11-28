import React, {useEffect, useState} from "react";
import PageToolbar from "./PageToolbar";
import MenuIcon from "@material-ui/icons/Menu";
import IconButton from "@material-ui/core/IconButton";
import Progress from "./Progress";
import Tab from "@material-ui/core/Tab";
import Hidden from "@material-ui/core/Hidden";
import Tabs from "@material-ui/core/Tabs";
import {
    Accordion,
    AccordionDetails,
    AccordionSummary,
    Avatar,
    Button,
    Container,
    Divider,
    Grid,
    Theme,
    Typography
} from "@material-ui/core";
import {logout} from "./AuthPage";
import {makeStyles} from "@material-ui/core/styles";
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import {api} from "../api/api";
import noAccountEmptyState from "./no-accounts-empty-state.svg";
import {permissionNameDescriptionMap} from "./permissionNameDescriptionMap";
import ConfirmationDialog from "./ConfirmationDialog";

const useStyles = makeStyles((theme: Theme) => ({
    indicator: {
        backgroundColor: '#fff',
    },
    expandIcon: {
        position: 'absolute', right: 32, top: 24, color: "#006580"
    }
}));

export default function Dashboard({authorizationServerURL, authorizationServerId, tenantId}) {
    const [isProgress, setProgress] = useState(true);
    const [consents, setConsents] = useState<any>([]);
    const [revokeDialog, setRevokeDialog] = useState<any>(null);
    const classes = useStyles();

    useEffect(() => {
        setProgress(true);
        api.getConsents()
            .then(res => setConsents(res.consents))
            .catch(err => console.log(err))
            .finally(() => setProgress(false));
    }, []);

    const handleRevoke = id => {
        setProgress(true);
        api.deleteConsent({id})
            .then(api.getConsents)
            .then(res => setConsents(res.consents))
            .catch(err => console.log(err))
            .finally(() => setProgress(false));
    }

    return (
        <div>
            <PageToolbar>
                <Hidden mdUp>
                    <IconButton edge="start" color="inherit" aria-label="menu">
                        <MenuIcon/>
                    </IconButton>
                </Hidden>
                <Hidden smDown>
                    <Tabs
                        value={'connected-applications'}
                        onChange={() => {
                        }}
                        classes={{
                            indicator: classes.indicator
                        }}
                        aria-label="menu tabs"
                        style={{height: 64}}
                    >
                        <Tab label="Get started" value={'get-started'}
                             style={{height: 64, textTransform: "none", minWidth: "unset"}}/>
                        <Tab label="APIs" value={'apis'}
                             style={{height: 64, textTransform: "none", minWidth: "unset"}}/>
                        <Tab label="Help Center" value={'help-center'}
                             style={{height: 64, textTransform: "none", minWidth: "unset"}}/>
                        <Tab label="Connected Applications" value={'connected-applications'}
                             style={{height: 64, textTransform: "none", minWidth: "unset"}}/>
                        <Tab label="Monitoring" value={'monitoring'}
                             style={{height: 64, textTransform: "none", minWidth: "unset"}}/>
                    </Tabs>
                </Hidden>
                <div style={{flex: 1}}/>
                <Button
                    style={{color: '#fff'}}
                    onClick={() => logout(authorizationServerURL, tenantId, authorizationServerId)}>Logout</Button>
            </PageToolbar>
            <div style={{marginTop: 64, position: "relative"}}>
                {isProgress && <Progress/>}

                {!isProgress && (
                    <>
                        {consents.length > 0 && (
                            <div style={{
                                background: "#F7F7F7",
                                height: 148,
                                display: 'flex',
                                alignItems: 'center'
                            }}>
                                <Container>
                                    <Typography variant={"h3"} color={"primary"}>Connected Applications</Typography>
                                </Container>
                            </div>
                        )}
                        <Container style={{marginTop: 64}}>
                            <Grid container justify={"center"}>
                                <Grid item xs={8}>
                                    {consents.length === 0 && (
                                        <div style={{textAlign: "center", marginTop: 64}}>
                                            <Typography variant={"h3"} style={{color: "#626576"}}>No connected
                                                accounts</Typography>
                                            <Typography style={{marginTop: 12, color: "#A0A3B5"}}>You haven’t connected
                                                any accounts yet
                                                to manage
                                                access</Typography>
                                            <img src={noAccountEmptyState} style={{marginTop: 64}} alt={"empty state"}/>
                                        </div>
                                    )}
                                    {consents.map(consent => (
                                        <Accordion key={consent.id}>
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
                                                        color: "#006580"
                                                    }}>{consent?.name[0]?.toUpperCase()}</Avatar>
                                                </div>
                                                <div style={{width: "100%"}}>
                                                    <div style={{display: "flex", alignItems: "center"}}>
                                                        <Typography variant="h4"
                                                                    style={{marginLeft: 24}}>{consent.name}</Typography>
                                                    </div>
                                                    <Divider style={{margin: "24px 0 24px 24px"}}/>
                                                    <div style={{
                                                        display: "flex",
                                                        justifyContent: "space-between",
                                                        marginLeft: 24
                                                    }}>
                                                        <div>
                                                            <Typography
                                                                style={{color: "#626576", fontSize: 16}}><strong>Consent
                                                                validation</strong></Typography>
                                                            <div style={{display: "flex", marginTop: 16}}>
                                                                <Typography>Granted on</Typography>
                                                                <div style={{
                                                                    background: "#D5E2E5",
                                                                    color: "#006580",
                                                                    fontSize: 12,
                                                                    padding: 4,
                                                                    fontWeight: 500,
                                                                    marginLeft: 8
                                                                }}
                                                                >{consent.creation_date_time.split('T')[0]}
                                                                </div>
                                                                <Typography style={{marginLeft: 36}}>Valid
                                                                    until</Typography>
                                                                <div style={{
                                                                    background: "#D5E2E5",
                                                                    color: "#006580",
                                                                    fontSize: 12,
                                                                    padding: 4,
                                                                    fontWeight: 500,
                                                                    marginLeft: 8
                                                                }}
                                                                >{consent.expiration_date_time.split('T')[0]}
                                                                </div>
                                                            </div>
                                                        </div>
                                                        <Button
                                                            style={{color: "#DC1B37", textTransform: 'none'}}
                                                            onClick={e => {
                                                                e.stopPropagation();
                                                                setRevokeDialog(consent);
                                                            }}
                                                        >Revoke Access
                                                        </Button>
                                                    </div>
                                                    <div style={{marginLeft: 24, marginTop: 32}}>
                                                        <Button
                                                            variant={"outlined"}
                                                            color={"primary"}
                                                            style={{textTransform: 'none'}}
                                                            onClick={e => {
                                                                e.stopPropagation()
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
                                                // paddingLeft: 88
                                            }}>
                                                {consent.account_ids.map(accountId => (
                                                    <div style={{paddingBottom: 32, borderLeft: "6px solid #006580",}}
                                                         key={accountId}>
                                                        <div style={{
                                                            background: "#D3E5EA",
                                                            height: 72,
                                                            color: "#006580",
                                                            paddingLeft: 88,
                                                            display: "flex",
                                                            alignItems: "center",
                                                            textTransform: 'uppercase'
                                                        }}>
                                                            <Typography>
                                                                <strong>Account ID: {accountId}</strong>
                                                            </Typography>
                                                        </div>
                                                        <div style={{paddingLeft: 88, color: "#006580"}}>
                                                            <Typography style={{
                                                                marginTop: 24,
                                                                borderBottom: "1px solid #ECECEC",
                                                                paddingBottom: 16,
                                                                marginBottom: 24,
                                                                textTransform: 'uppercase'
                                                            }}><strong>Granted permissions:</strong></Typography>
                                                            <Grid container spacing={3}>
                                                                {consent.permissions.map(permission => (
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
                                                        </div>
                                                    </div>
                                                ))}
                                            </AccordionDetails>
                                        </Accordion>
                                    ))}
                                </Grid>
                            </Grid>

                        </Container>
                    </>
                )}
            </div>
            {revokeDialog && (
                <ConfirmationDialog
                    title="Revoke application access"
                    content="Are you sure you want to revoke access to your accounts for this application?"
                    confirmText="Yes, Revoke Access"
                    warningItems={[
                        'Your connected application will not be able to access your bank accounts.',
                        'You will have to authorize the application again if you would like to use it with your bank in the future.'
                    ]}
                    onCancel={() => setRevokeDialog(null)}
                    onConfirm={() => {
                        handleRevoke(revokeDialog.id);
                        setRevokeDialog(null);
                    }}
                />
            )}
        </div>
    )
};
