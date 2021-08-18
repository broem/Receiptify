<template>
  <div id="landing">
    <div class="create_account">
      <v-col cols="auto">
        <v-dialog transition="dialog-bottom-transition" max-width="600">
          <template v-slot:activator="{ on, attrs }">
            <v-chip
              style="width:120px; height:40px; text-align:center;"
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
                      rules="required|name_length"
                      debounce="1000"
                      v-slot="{ errors }"
                    >
                      <v-text-field
                        v-model="name"
                        hide-details="auto"
                        label="Name"
                      ></v-text-field>
                      <span id="error">{{ errors[0] }}</span>
                    </ValidationProvider>
                    <ValidationProvider
                      rules="required|email"
                      debounce="1000"
                      v-slot="{ errors }"
                    >
                      <v-text-field
                        hide-details="auto"
                        v-model="email"
                        label="Email"
                      ></v-text-field>
                      <span id="error">{{ errors[0] }}</span>
                    </ValidationProvider>
                    <ValidationProvider
                      rules="required|password_length|password_strength|confirmed:confirmation"
                      debounce="1000"
                      v-slot="{ errors }"
                    >
                      <v-text-field
                        type="password"
                        hide-details="auto"
                        v-model="password"
                        label="Password"
                      ></v-text-field>
                      <span id="error">{{ errors[0] }}</span>
                    </ValidationProvider>
                    <ValidationProvider
                      rules="required"
                      vid="confirmation"
                      debounce="1000"
                      v-slot="{ errors }"
                    >
                      <v-text-field
                        type="password"
                        hide-details="auto"
                        v-model="confirm_password"
                        label="Confirm Password"
                      ></v-text-field>
                      <span id="error">{{ errors[0] }}</span>
                    </ValidationProvider>

                    <v-btn :disabled="invalid" type="submit" value="Submit">
                      Submit
                    </v-btn>

                    <v-btn color="error" class="mr-4" @click="reset">
                      Reset Form
                    </v-btn>

                    <v-btn color="warning" @click="resetValidation">
                      Reset Validation
                    </v-btn>
                  </v-form></ValidationObserver
                >
              </v-card-text>
              <v-card-actions class="justify-end">
                <v-btn text @click="dialog.value = false">Close</v-btn>
              </v-card-actions>
            </v-card>
          </template>
        </v-dialog>
      </v-col>
    </div>
    <div class="digitize_label">
      <img class="d_label" src="@/assets/digitizeLabel.png" />
    </div>
    <div class="digitize">
      <router-link to="/digitize"
        ><img src="@/assets/digitize.png" />
      </router-link>
    </div>

    <div class="content">
      <event-hub></event-hub>
      <router-view></router-view>
    </div>
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
    return (value.length < 10 && value.length > 1) || "Name is invalid.";
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
  name: "landing",

  data: () => {
    return {
      validated: false,
    };
  },
  components: {
    ValidationProvider,
    ValidationObserver,
  },
  methods: {
    validate() {
      this.$refs.form.validate();
    },
    reset() {
      this.$refs.form.reset();
    },
    resetValidation() {
      this.$refs.form.resetValidation();
    },
    async submit() {
      const isValid = await this.$refs.observer.validate();
      if (isValid) {
        console.log("validated! and account added!"); //logic for adding account information goes here.
      } else {
        alert("Data isn't valid");
      }
    },
  },
};
</script>

<style>
#landing {
  position: relative;
  width: 100%;
  height: 100vh;
}

.account_registration_header {
  background-color: rgba(47, 44, 54, 0.397);
}

@media (max-width: 599px) {
  .create_account {
    position: absolute;
    left: 60%;
    top: 50%;
    bottom: 50%;
    width: 140px;
    height: 50px;
  }
  .digitize {
    position: absolute;
    object-fit: contain;
    width: 100%;
    left: 55%;
    bottom: 0%;
  }
  .digitize_label {
    position: absolute;
    object-fit: contain;
    width: auto;
    left: 15%;
    bottom: 7%;
  }

  .d_label {
    width: 200px !important;
  }
}
@media (min-width: 600px) {
  .create_account {
    position: absolute;
    left: 60%;
    top: 50%;
    bottom: 50%;
    width: 140px;
    height: 50px;
  }

  .digitize {
    position: absolute;
    object-fit: contain;
    width: 100%;
    left: 60%;
    bottom: 2%;
  }
  .digitize_label {
    position: absolute;
    object-fit: contain;
    width: auto;
    left: 15%;
    bottom: 7%;
  }

  .d_label {
    width: 200px !important;
  }
}

@media (min-width: 1100px) {
  .create_account {
    position: absolute;
    left: 80%;
    top: 50%;
    bottom: 50%;
    width: 140px;
    height: 50px;
  }
  .digitize {
    position: absolute;
    object-fit: contain;
    width: 100%;
    left: 50%;
    bottom: 2%;
  }
  .digitize_label {
    position: absolute;
    object-fit: contain;
    width: auto;
    left: 15%;
    bottom: 7%;
  }

  .d_label {
    width: 400px !important;
  }
}
</style>
