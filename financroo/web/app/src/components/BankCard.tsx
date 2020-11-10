import React from "react";
import {Theme} from "@material-ui/core";
import {makeStyles} from "@material-ui/core/styles";
import apexfinancialIcon from "../assets/banks/apexfinancial-icon.svg";
import Typography from "@material-ui/core/Typography";
import IconButton from "@material-ui/core/IconButton";
import {Plus} from "react-feather";
import Card from "@material-ui/core/Card";
import Checkbox from "@material-ui/core/Checkbox";

const useStyles = makeStyles((theme: Theme) => ({}));

export default ({account, selected, style = {}, onToggleSelect}) => {
  const classes = useStyles();

  return (
    <Card style={style}>
      <div style={{padding: 20, display: 'flex', alignItems: 'center'}}>
        <div style={{
          background: '#FCFCFF',
          borderRadius: '50%',
          width: 52,
          height: 52,
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center',
          boxShadow: '0px 0.574468px 0.574468px rgba(0, 0, 0, 0.08), 0px 0px 0.574468px rgba(0, 0, 0, 0.31)'
        }}>
          <img src={apexfinancialIcon} style={{width: 24, height: 24}}/>
        </div>
        <div style={{marginLeft: 24}}>
          <Typography>Apex Financial</Typography>
          <Typography style={{background: 'rgba(54, 198, 175, 0.08)', color: '#36C6AF', fontSize: 14, padding: 2, marginTop: 4}}>2 accounts
            synced</Typography>
        </div>
      </div>
      <div style={{
        height: 62,
        background: '#36C6AF',
        color: '#fff',
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'space-between',
        padding: '0 20px'
      }}>
        <div style={{display: 'flex', alignItems: 'center'}}>
          <Checkbox
            checked={selected}
            onChange={onToggleSelect}
            color={'primary'}
            style={{color: selected && '#fff'}}
            inputProps={{'aria-label': 'primary checkbox'}}
          />
          <div style={{marginLeft: 12}}>
            <Typography>Credit account</Typography>
            <Typography>**** ***** **** 4099</Typography>
          </div>
        </div>
        <div>
          <Typography>€ 99920.20</Typography>
        </div>
      </div>
      <div style={{height: 62, display: 'flex', alignItems: 'center', justifyContent: 'space-between', padding: '0 20px'}}>
        <div style={{display: 'flex', alignItems: 'center'}}>
          <Checkbox
            checked={false}
            onChange={onToggleSelect}
            color={'primary'}
            inputProps={{'aria-label': 'primary checkbox'}}
          />
          <div style={{marginLeft: 12}}>
            <Typography>Savings account</Typography>
            <Typography>**** ***** **** 1134</Typography>
          </div>
        </div>
        <div>
          <Typography>€ 20550.20</Typography>
        </div>
      </div>
      <div style={{
        height: 52,
        padding: '0 21px',
        background: 'rgba(54, 198, 175, 0.08)',
        color: '#36C6AF',
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'space-between'
      }}>
        <Typography>Add new account</Typography>
        <IconButton>
          <Plus style={{color: '#36C6AF'}}/>
        </IconButton>
      </div>
    </Card>
  )
};
