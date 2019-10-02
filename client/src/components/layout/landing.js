import React, { Component } from 'react';
import axios from 'axios';

class Landing extends Component {
    constructor() {
        super();
        this.state = {
            counter: 0
        };
    }

    componentDidMount() {
        axios.defaults.headers.common['CLIENT_IP'] = "123.1.1.12";
        axios.get('/api/counter')
            .then(res => {
                console.log(res.data.counter);
                this.setState({ counter: res.data.counter })
            })
            .catch(err => {
                console.log(err);
            });
    }

    render = () => {
        return (
            <div className="landingPage">
                <h1>GAWDS TASK</h1>
                <h1>{this.state.counter} user(s) visited this page</h1>
            </div>
        )
    }
}

export default Landing;