const Migrations = artifacts.require("Migrations");
const Exchange = artifacts.require("Exchange");
const ExchangeData = artifacts.require("ExchangeData");
// const TestToken = artifacts.require("TestToken");

module.exports = function(deployer) {
  deployer.deploy(Migrations);
  deployer.deploy(Exchange);
  deployer.deploy(ExchangeData);

  // deployer.deploy(TestToken,2,1000000000,"BDE");
  // deployer.deploy(TestToken,18,1000000000,"DDD");
  // deployer.deploy(TestToken,18,1000000000,"EEE");
};
