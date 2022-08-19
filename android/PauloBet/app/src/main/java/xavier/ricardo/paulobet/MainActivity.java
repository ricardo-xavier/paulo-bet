package xavier.ricardo.paulobet;

import androidx.appcompat.app.AppCompatActivity;

import android.content.Context;
import android.content.Intent;
import android.content.SharedPreferences;
import android.os.Bundle;
import android.view.View;
import android.widget.CheckBox;
import android.widget.EditText;
import android.widget.Toast;

import com.google.gson.Gson;

import java.io.IOException;

import okhttp3.Response;
import xavier.ricardo.paulobet.model.LoginResponse;
import xavier.ricardo.paulobet.tasks.LoginTask;

public class MainActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        loadPreferences();
    }

    public void login(View v) {
        EditText etLogin = (EditText) findViewById(R.id.etLogin);
        String user = etLogin.getText().toString().trim();
        if (user.isEmpty()) {
            Toast.makeText(this, "Usuário não preenchido", Toast.LENGTH_LONG).show();
            etLogin.requestFocus();
            return;
        }
        EditText tpPassword = (EditText) findViewById(R.id.tpPassword);
        String password = tpPassword.getText().toString().trim();
        new LoginTask(this, user, password).execute();
    }

    public void onTaskResponse(Response response) throws IOException {
        if (response.code() == 200) {
            Gson gson = new Gson();
            LoginResponse loginResponse = gson.fromJson(response.body().string(), LoginResponse.class);
            savePreferences();
            EditText etLogin = (EditText) findViewById(R.id.etLogin);
            String user = etLogin.getText().toString().trim();
            Intent intent = new Intent(this, SelectLeagueActivity.class);
            intent.putExtra("user", user);
            intent.putExtra("token", loginResponse.getToken());
            startActivity(intent);
        } else {
            Toast.makeText(this, "ERRO status " + response.code(), Toast.LENGTH_LONG).show();
        }
    }

    private void loadPreferences() {
        SharedPreferences shared = getSharedPreferences("paulobet", Context.MODE_PRIVATE);
        boolean save = shared.getBoolean("save", false);
        if (save) {
            CheckBox cbSave = (CheckBox) findViewById(R.id.cbSave);
            cbSave.setChecked(true);

            String user = shared.getString("user", "");
            if (!user.isEmpty()) {
                EditText etLogin = (EditText) findViewById(R.id.etLogin);
                etLogin.setText(user);
            }
            String password = shared.getString("password", "");
            if (!password.isEmpty()) {
                EditText tpPassword = (EditText) findViewById(R.id.tpPassword);
                tpPassword.setText(password);
            }
        }
    }

    private void savePreferences() {
        SharedPreferences shared = getSharedPreferences("paulobet", Context.MODE_PRIVATE);
        CheckBox cbSave = (CheckBox) findViewById(R.id.cbSave);
        if (!cbSave.isChecked()) {
            return;
        }
        EditText etLogin = (EditText) findViewById(R.id.etLogin);
        String user = etLogin.getText().toString().trim();
        EditText tpPassword = (EditText) findViewById(R.id.tpPassword);
        String password = tpPassword.getText().toString().trim();
        SharedPreferences.Editor editor = shared.edit();
        editor.putString("user", user);
        editor.putString("password", password);
        editor.putBoolean("save", true);
        editor.commit();
    }
}