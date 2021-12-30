function mySettings(props) {
  return (
    <Page>
      <Section
        title={<Text bold align="center">ECP Address</Text>}>
        <TextInput
          label="ECP Host"
          settingsKey="ecphost"
          placeholder="Ex: 192.168.0.211" />
        <TextInput
          label="ECP Port"
          settingsKey="ecpport"
          placeholder="Ex: 8060" />
      </Section>
    </Page>
  );
}

registerSettingsPage(mySettings);
