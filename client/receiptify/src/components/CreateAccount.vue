<template>
  <div id="createAccount">
    <v-col cols="auto">
      <v-dialog transition="dialog-bottom-transition" max-width="600">
        <template v-slot:activator="{ on, attrs }">
          <v-chip
            style="width:170px; height:40px; text-align:center;"
            color="blue"
            v-bind="attrs"
            v-on="on"
            outlined
            >Create Account >
          </v-chip>
        </template>
        <template v-slot:default="dialog">
          <v-card>
            <v-toolbar class="account_registration_header" dark
              >Lets get an account setup.</v-toolbar
            >
            <v-card-text>
              <ValidationObserver ref="observer" v-slot="{ invalid }">
                <v-form
                  ref="form"
                  v-model="valid"
                  lazy-validation
                  v-on:submit.prevent="submit()"
                >
                  <ValidationProvider
                    rules="required|name_length|name_chars"
                    debounce="100"
                    v-slot="{ errors }"
                  >
                    <v-text-field
                      class="field"
                      v-model="last_name"
                      hide-details="auto"
                      label="Last Name"
                    ></v-text-field>
                    <div class="error_box" id="error">{{ errors[0] }}</div>
                  </ValidationProvider>
                  <ValidationProvider
                    rules="required|name_length|name_chars"
                    debounce="100"
                    v-slot="{ errors }"
                  >
                    <v-text-field
                      class="field"
                      v-model="first_name"
                      hide-details="auto"
                      label="First Name"
                    ></v-text-field>
                    <div class="error_box" id="error">{{ errors[0] }}</div>
                  </ValidationProvider>

                  <ValidationProvider
                    rules="required|email"
                    debounce="100"
                    v-slot="{ errors }"
                  >
                    <v-text-field
                      class="field"
                      v-model="email"
                      hide-details="auto"
                      label="Email"
                    ></v-text-field>
                    <div class="error_box" id="error">{{ errors[0] }}</div>
                  </ValidationProvider>
                  <ValidationProvider
                    rules="required|password_length|password_strength|confirmed:confirmation"
                    debounce="100"
                    v-slot="{ errors }"
                  >
                    <v-text-field
                      class="field"
                      type="password"
                      v-model="password"
                      hide-details="auto"
                      label="Password"
                    ></v-text-field>
                    <div class="error_box" id="error">{{ errors[0] }}</div>
                  </ValidationProvider>
                  <ValidationProvider
                    rules="required"
                    vid="confirmation"
                    debounce="100"
                    v-slot="{ errors }"
                  >
                    <v-text-field
                      class="field"
                      type="password"
                      hide-details="auto"
                      v-model="confirm_password"
                      label="Confirm Password"
                    ></v-text-field>
                    <div class="error_box" id="error">{{ errors[0] }}</div>
                  </ValidationProvider>
                  <div class="button_row">
                    <v-btn
                      class="global_button"
                      :disabled="invalid"
                      type="submit"
                      value="Submit"
                    >
                      Create
                    </v-btn>

                    <v-btn
                      class="global_button mr-4"
                      color="error"
                      @click="reset"
                    >
                      Reset Form
                    </v-btn>
                    <v-btn
                      class="global_button"
                      text
                      @click="dialog.value = false"
                      >Close</v-btn
                    >
                  </div>
                </v-form></ValidationObserver
              >
            </v-card-text>
          </v-card>
        </template>
      </v-dialog>
    </v-col>
  </div>
</template>

<script>
import { ValidationProvider, ValidationObserver } from "vee-validate";
import { extend } from "vee-validate";
import { required, confirmed } from "vee-validate/dist/rules";

extend("confirmed", {
  ...confirmed,
  message: "Passwords needs to match.",
});

extend("required", {
  ...required,
  message: "This field is required",
});

extend("name_length", {
  validate: (value) => {
    return (value.length < 15 && value.length > 1) || "Name is invalid.";
  },
});

extend("name_chars", {
  validate: (value) => {
    return (
      (/[^a-zA-Z0-9]+/.test(value) == false && /[0-9]+/.test(value) == false) ||
      "Your using invalid characters."
    );
  },
});

extend("email", {
  validate: (value) => {
    return (
      /^\w+([.-]?\w+)*@\w+([.-]?\w+)*(\.\w{2,3})+$/.test(value) ||
      "E-mail must be valid"
    );
  },
});

extend("password_length", {
  validate: (value) => {
    return value.length >= 8 || "Use 8 characters or more for your password";
  },
});

extend("password_strength", {
  validate: (value) => {
    return (
      (/[a-z]+/.test(value) &&
        /[A-Z]+/.test(value) &&
        /\d+/.test(value) &&
        /[^a-zA-Z0-9]+/.test(value)) ||
      "Please choose a stronger password. Try a mix of letters, numbers, and symbols."
    );
  },
});
export default {
  name: "createAccount",
  components: {
    ValidationProvider,
    ValidationObserver,
  },
  methods: {
    reset() {
      this.$refs.form.reset();
    },
    async submit() {
      const isValid = await this.$refs.observer.validate();
      if (isValid) {
        alert("Account was successfully Created."); //account is ready to be added.
        this.reset();
      } else {
        alert("Data isn't valid");
      }
    },
  },
};
</script>

<style>
.account_registration_header {
  background-color: rgba(47, 44, 54, 0.397);
  margin-bottom: 30px;
}

.field {
  position: relative;
  margin: 5px;
  padding: 10px;
  height: 50px;
}

.error_box {
  height: 30px;
  color: red;
}

.button_row {
  padding: 1rem;
  position: relative;
  display: flex;
  align-items: center;
  height: 40px;
  width: auto;
}

.global_button {
  margin-left: 5px;
  margin-right: 5px;
  position: relative;
}
</style>
