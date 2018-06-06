import React, { Component } from "react";
import { View, Text, FlatList } from "react-native";
import { SearchBar, Header } from "react-native-elements";
import Ionicons from 'react-native-vector-icons/Ionicons';
import { createBottomTabNavigator } from 'react-navigation';

class App extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      data: [],
      error: null
    };

    this.getDialogsFromApiAsync();
  }
  componentDidMount() {
    this.getDialogsFromApiAsync();
  }

  getDialogsFromApiAsync = () => {
    return fetch("http://192.168.0.102:8080/dialog/0/10")
      .then(res => res.json())
      .then(response => {
        this.setState({ data: response });
      })
      .catch(error => {
        console.log(error);
        this.setState({ error });
      });
  };

  renderSeparator = () => {
    return (
      <View
        style={{
          height: 1,
          width: "100%",
          backgroundColor: "#CED0CE",
        }}
      />
    );
  };
  
  _renderItem = ({item}) => (
        <View style={{ flexDirection:'row'}}>
          <Text style={{minHeight: 30, backgroundColor: 'black',color:'white',textAlign:'justify',padding:10,flex: 1, flexWrap: 'wrap'}} >{item.dialog}</Text>
        </View>
  );
  render() {
    return (
      <View>
        <Header
          leftComponent={{ icon: "menu", color: "#fff" }}
          centerComponent={{ text: "Picas", style: { color: "#fff" } }}
          rightComponent={{ icon: "home", color: "#fff" }}
        />
        <SearchBar
          showLoading
          platform="ios"
          cancelButtonTitle="Cancel"
          placeholder="Search for a picas"
          onFocus={this.getDialogsFromApiAsync}
        />
        <View>
          <FlatList
            data={this.state.data}
            keyExtractor={item => item._id}
            renderItem={this._renderItem}
            ItemSeparatorComponent={this.renderSeparator}
          />
        </View>
      </View>
    );
  }
}

class SettingsScreen extends React.Component {
  render() {
    return (
      <View style={{ flex: 1, justifyContent: 'center', alignItems: 'center' }}>
        <Text>Settings!</Text>
      </View>
    );
  }
}

class ProfileScreen extends React.Component {
  render() {
    return (
      <View style={{ flex: 1, justifyContent: 'center', alignItems: 'center' }}>
        <Text>Profile</Text>
      </View>
    );
  }
}

export default createBottomTabNavigator({
  Home: App,
  Settings: SettingsScreen,
  MyProfile: ProfileScreen,
},
{
  navigationOptions: ({ navigation }) => ({
    tabBarIcon: ({ focused, tintColor }) => {
      const { routeName } = navigation.state;
      let iconName;
      if (routeName === 'Home') {
        iconName = `ios-book${focused ? '' : '-outline'}`;
      } else if (routeName === 'Settings') {
        iconName = `ios-options${focused ? '' : '-outline'}`;
      }else if (routeName === 'MyProfile') {
          iconName = `ios-people${focused ? '' : '-outline'}`;
      }

      // You can return any component that you like here! We usually use an
      // icon component from react-native-vector-icons
      return <Ionicons name={iconName} size={25} color={tintColor} />;
    },
  }),
  tabBarOptions: {
    activeTintColor: 'tomato',
    inactiveTintColor: 'gray',
  },
}
);