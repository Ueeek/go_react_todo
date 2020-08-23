import React, { Component } from "react";
import axios from "axios";
import {Button, Card, Header, Form, Input, Icon, Dropdown} from "semantic-ui-react";

let endpoint = "http://localhost:8080";

class ToDoList extends Component {
  constructor(props) {
    super(props);

    this.state = {
      task: "",
      items: [],
      sort_by:"date",
      sort_order:"acs",
      show_option:"all"
    };
  }

  componentDidMount() {
    this.getTask();
  }

  onChange = event => {
    this.setState({
      [event.target.name]: event.target.value
    });
  };

  onSubmit = () => {
    let { task } = this.state;
    console.log("pRINTING task", this.state.task);
    if (task) {
      axios
        .post(
          endpoint + "/api/task",
          {
            task
          },
          {
            headers: {
              "Content-Type": "application/x-www-form-urlencoded"
            }
          }
        )
        .then(res => {
          this.getTask();
          this.setState({
            task: ""
          });
          console.log(res);
        });
    }
  };

  getTask = () => {
    axios.get(endpoint + "/api/task",{
        params:{
            sort_by: this.state.sort_by,
            sort_order: this.state.sort_order,
            show_option:this.state.show_option,
        }}).then(res => {
      console.log(res);
      if (res.data) {
        this.setState({
          items: res.data.map(item => {
            let color = "yellow";

            if (item.status) {
              color = "green";
            }
            return (
              <Card key={item._id} color={color} fluid>
                <Card.Content>
                  <Card.Header textAlign="left">
                    <div style={{ wordWrap: "break-word" }}>{item.task}</div>
                  </Card.Header>
                  <Card.Meta>
                    <div>{item.date}</div>
                  </Card.Meta>

                  <Card.Meta textAlign="right">
                    <Icon
                      name="check circle"
                      color="green"
                      onClick={() => this.updateTask(item._id)}
                    />
                    <span style={{ paddingRight: 10 }}>Done</span>
                    <Icon
                      name="undo"
                      color="yellow"
                      onClick={() => this.undoTask(item._id)}
                    />
                    <span style={{ paddingRight: 10 }}>Undo</span>
                    <Icon
                      name="delete"
                      color="red"
                      onClick={() => this.deleteTask(item._id)}
                    />
                    <span style={{ paddingRight: 10 }}>Delete</span>
                  </Card.Meta>
                </Card.Content>
              </Card>
            );
          })
        });
      } else {
        this.setState({
          items: []
        });
      }
    });
  };

  updateTask = id => {
    axios
      .put(endpoint + "/api/task/" + id, {
        headers: {
          "Content-Type": "application/x-www-form-urlencoded"
        }
      })
      .then(res => {
        console.log(res);
        this.getTask();
      });
  };

  undoTask = id => {
    axios
      .put(endpoint + "/api/undoTask/" + id, {
        headers: {
          "Content-Type": "application/x-www-form-urlencoded"
        }
      })
      .then(res => {
        console.log(res);
        this.getTask();
      });
  };

  deleteTask = id => {
    axios
      .delete(endpoint + "/api/deleteTask/" + id, {
        headers: {
          "Content-Type": "application/x-www-form-urlencoded"
        }
      })
      .then(res => {
        console.log(res);
        this.getTask();
      });
  };

  deleteAllTask = () =>{
      axios
        .delete(endpoint+"/api/deleteAllTask",{
            headers:{
                "Content-Type":"application/x-www-form-urlencoded"
            }
        })
        .then(res=>{
            console.log(res);
            this.getTask();
        });
  };
  deleteDoneTask = () =>{
      axios
        .delete(endpoint+"/api/deleteDoneTask",{
            headers:{
                "Content-Type":"application/x-www-form-urlencoded"
            }
        })
        .then(res=>{
            console.log(res);
            this.getTask();
        });
  };

  onChangeSortBy = async(e,{value}) =>{
      await this.setState({sort_by:value})
      await this.getTask();
  }
  onChangeSortOrder = async(e,{value}) =>{
      await this.setState({sort_order:value})
      await this.getTask();
  }

  onChangeShowOption = async(e,{value})=>{
      await this.setState({show_option:value})
      await this.getTask();
  }

  render() {
    return (
      <div>
        <div className="row">
          <Header className="header" as="h2">
            TO DO LIST
          </Header>
        </div>
        <div className="row">
            <Button onClick={this.deleteDoneTask}> Delete Done Tasks </Button>
            <Button onClick={this.deleteAllTask}> Delete ALL Tasks </Button>
        </div>
        <div className="row">
            <div className="select">
                <Dropdown
                    placeholder="show option"
                    name="show option"
                    selection
                    onChange={this.onChangeShowOption}
                    options={[{text:"All",value:"all",key:"all"},
                              {text:"Done",value:"done",key:"done"},
                              {text:"Yet",value:"yet",key:"yet"}]}
                />
            </div>
            <div className="select">
                <Dropdown
                    placeholder="sort by"
                    onChange={this.onChangeSortBy}
                    name="sort by"
                    selection
                    options={[{text:"status",value:"status",key:"status"},
                              {text:"date",value:"date",key:"date"},
                              {text:"task",value:"task",key:"task"}]}
                />
            </div>
            <div className="select">
                <Dropdown
                    placeholder="sort order"
                    onChange={this.onChangeSortOrder}
                    name="sort order"
                    selection
                    options={[{text:"acsending",value:"acs",key:"acs"},
                              {text:"decsenging",value:"dec",key:"dec"}]}
                />
            </div>
        </div>

        <div className="row">
          <Form onSubmit={this.onSubmit}>
            <Input
              type="text"
              name="task"
              onChange={this.onChange}
              value={this.state.task}
              fluid
              placeholder="Create Task"
            />
            {/* <Button >Create Task</Button> */}
          </Form>
        </div>
        <div className="row">
          <Card.Group>{this.state.items}</Card.Group>
        </div>
      </div>
    );
  }
}

export default ToDoList;
